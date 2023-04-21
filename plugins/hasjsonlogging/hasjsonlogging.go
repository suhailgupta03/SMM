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

func (j *HasJSONLogging) Check(logGroupLogStream string) types.MaturityCheck {
	inputSplit := strings.Split(logGroupLogStream, "_")
	if len(inputSplit) < 2 {
		return types.MaturityValue0
	}
	awsLogGroup := inputSplit[0]
	awsLogStream := inputSplit[1]
	aws := util.AWSInit()
	logs, err := aws.GetLogs(awsLogGroup, awsLogStream)
	if err != nil {
		return types.MaturityValue0
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
