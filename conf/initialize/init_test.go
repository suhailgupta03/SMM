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
	assert.False(t, envVars.ScanAllGitHub)

	os.Setenv("SCAN_ALL_GITHUB_REPOS", "true")
	envVars = getFromEnv()
	assert.True(t, envVars.ScanAllGitHub)

	// Revert the old val
	os.Setenv("STAGE", prevStageVal)
	os.Setenv("SCAN_ALL_GITHUB_REPOS", "false")
}
