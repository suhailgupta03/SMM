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
	envVars := getFromEnv(nil)
	assert.Nil(t, envVars.Test)
	assert.False(t, envVars.ScanAllGitHub)

	os.Setenv("SCAN_ALL_GITHUB_REPOS", "true")
	envVars = getFromEnv(nil)
	assert.True(t, envVars.ScanAllGitHub)

	// Revert the old val
	os.Setenv("STAGE", prevStageVal)
	os.Setenv("SCAN_ALL_GITHUB_REPOS", "false")
}

func TestGetAppConstants2(t *testing.T) {

	os.Args = []string{"...", "repo", "-yml=foobar.yml", "github", "-token=TOKEN", "-owner=OWNER"}
	f := initFlags()
	assert.Equal(t, "foobar.yml", *f.YAMLFile)
	assert.Equal(t, "TOKEN", *f.GitHubToken)
	assert.Equal(t, "OWNER", *f.GitHubOwner)

	os.Args = []string{}
	f = initFlags()
	assert.Equal(t, "repo-details.yml", *f.YAMLFile)
	assert.Equal(t, "", *f.GitHubToken)
	assert.Equal(t, "", *f.GitHubOwner)

	prevToken := os.Getenv("TOKEN")
	prevOwner := os.Getenv("OWNER")
	os.Setenv("TOKEN", "FOOBAR_TOKEN")
	os.Setenv("OWNER", "FOOBAR_OWNER")
	os.Args = []string{"...", "github", "-token=APP_TOKEN", "-owner=FOOBAR_OWNER"}
	f = initFlags()
	details := getFromEnv(&f)
	assert.Equal(t, "APP_TOKEN", details.GitHubToken)
	assert.Equal(t, "FOOBAR_OWNER", details.GitHubOwner)

	// Set back the values
	os.Setenv("TOKEN", prevToken)
	os.Setenv("OWNER", prevOwner)
}
