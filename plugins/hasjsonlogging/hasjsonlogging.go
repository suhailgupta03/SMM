package main

import (
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"encoding/json"
	"strings"
)

type HasJSONLogging struct {
}

// hasJSONStrings returns true if there is at least 1 log
// in the JSON format
func hasJSONStrings(list []*string) bool {
	var js json.RawMessage
	totalValidStringsFound := 0
	for _, log := range list {
		err := json.Unmarshal([]byte(*log), &js)
		if err == nil {
			totalValidStringsFound += 1
			// Can also simply break and return !!
		}
	}

	if totalValidStringsFound > 0 {
		return true
	}
	return false
}

// Check accepts logGroupLogStream and an optional slice of string pointers
// The first index of opts is the CWLQuery
// If the CWL query has not been passed, will use the passed log group and the log
// stream to fetch and parse the logs
func (j *HasJSONLogging) Check(logGroupLogStream string, opts ...*string) types.MaturityCheck {
	inputSplit := strings.Split(logGroupLogStream, "_")
	if len(inputSplit) < 2 {
		return types.MaturityValue0
	}
	awsLogGroup := inputSplit[0]
	awsLogStream := inputSplit[1]

	var cwlQuery *string
	if len(opts) > 0 {
		cwlQuery = opts[0]
	}
	aws := util.AWSInit()

	logs := make([]*string, 0)

	if cwlQuery != nil {
		// If the CWL Query has been passed, prioritize query to
		// fetch the logs
		queryLogs, queryErr := aws.GetLogsUsingQuery(awsLogGroup, *cwlQuery)
		if queryErr != nil {
			return types.MaturityValue0
		}
		logs = queryLogs
	} else {
		// If the CWL query has not been passed, use the log group
		// and log stream to fetch the logs
		logStream, err := aws.GetLogs(awsLogGroup, awsLogStream)
		if err != nil {
			return types.MaturityValue0
		}
		logs = logStream
	}

	if hasJSONStrings(logs) {
		return types.MaturityValue2
	}

	return types.MaturityValue1
}

func (j *HasJSONLogging) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type:    types.MaturityObservability,
		Name:    "Has JSON logging",
		EcrType: false,
	}
}

var Check HasJSONLogging
