package initialize

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetAppConstants(t *testing.T) {
	isTest := isStageTest("prod")
	assert.False(t, isTest)

	prevStageVal := os.Getenv("STAGE")
	os.Setenv("STAGE", "prod")
	envVars := getFromEnv()
	assert.Nil(t, envVars.Test)
	assert.Nil(t, envVars.MaturityRepoDetails)

	// Revert the old val
	os.Setenv("STAGE", prevStageVal)
}
