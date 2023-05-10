package main

import (
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"strings"
)

type HasLogging struct {
}

// Check takes in '-' separate log group followed by the log stream
// example: myloggroup_mylogstream would mean the log group as
// myloggroup and log stream as mylogstream
func (l *HasLogging) Check(logGroupLogStream string, opts ...*string) types.MaturityCheck {
	aws := util.AWSInit()
	inputSplit := strings.Split(logGroupLogStream, "_")
	if len(inputSplit) < 2 {
		return types.MaturityValue0
	}
	awsLogGroup := inputSplit[0]
	awsLogStream := inputSplit[1]
	exists := aws.DoesLogStreamExist(awsLogGroup, awsLogStream)
	if exists {
		return types.MaturityValue2
	}
	return types.MaturityValue1
}

func (l *HasLogging) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type:    types.MaturityObservability,
		Name:    "Has (some) logging",
		EcrType: false,
	}
}

var Check HasLogging
