package util

import (
	"cuddly-eureka-/conf/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAWS_DoesLogGroupExist(t *testing.T) {
	aws := AWSInit()
	app := initialize.GetAppConstants()
	exists := aws.DoesLogGroupExist(app.Test.AWS.LogGroup)
	assert.True(t, exists)

	exists = aws.DoesLogGroupExist("random-group-name")
	assert.False(t, exists)
}

func TestAWS_DoesLogStreamExist(t *testing.T) {
	aws := AWSInit()
	app := initialize.GetAppConstants()
	exists := aws.DoesLogStreamExist(app.Test.AWS.LogGroup, app.Test.AWS.LogStream)
	assert.True(t, exists)

	exists = aws.DoesLogStreamExist("random-group-name", "random-stream-name")
	assert.False(t, exists)
}

func TestAWS_GetLogs(t *testing.T) {
	aws := AWSInit()
	app := initialize.GetAppConstants()
	logs, err := aws.GetLogs(app.Test.AWS.LogGroup, app.Test.AWS.LogStream)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
}

func TestAWS_GetLogsUsingQuery(t *testing.T) {
	aws := AWSInit()
	app := initialize.GetAppConstants()
	query := `sort @timestamp desc
| limit 20`

	_, err := aws.GetLogsUsingQuery(app.Test.AWS.LogGroup, query)
	assert.Nil(t, err)

	_, err = aws.GetLogsUsingQuery("...", "..")
	assert.NotNil(t, err)
}
