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
