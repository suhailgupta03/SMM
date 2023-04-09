package main

import (
	"cuddly-eureka-/conf/initialize"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigurationLoad(t *testing.T) {
	prevStageVal := os.Getenv("STAGE")

	appConstants = initialize.GetAppConstants()
	assert.Equal(t, "test", appConstants.Stage, "Stage name must be test")
	assert.NotEmpty(t, appConstants.GitHubOwner, "Tests if the github owner has been loaded")
	assert.NotEmpty(t, appConstants.GitHubToken, "Tests if the github token has been loaded")
	assert.NotNil(t, appConstants.Test, "Test configuration must not be nil")
	assert.NotNil(t, appConstants.MaturityRepoDetails)

	// Revert the old val
	os.Setenv("STAGE", prevStageVal)
}

func TestGetRepos(t *testing.T) {
	pre := os.Getenv("SCAN_ALL_GITHUB_REPOS")

	os.Setenv("SCAN_ALL_GITHUB_REPOS", "true")
	appConstants = initialize.GetAppConstants()
	repos := getRepos(appConstants.GitHubToken, appConstants.GitHubOwner)
	assert.GreaterOrEqual(t, len(repos), 0)

	os.Setenv("SCAN_ALL_GITHUB_REPOS", "false")
	appConstants = initialize.GetAppConstants()
	repos = getRepos(appConstants.GitHubToken, appConstants.GitHubOwner)
	assert.GreaterOrEqual(t, len(repos), 0)

	repoYML := `name: Repository Details
# Inside repository
# name is mandatory
# ecr is optional
repository:
  - name: virality
    ecr: xxxx.dkr.ecr.us-east-1.amazonaws.com/ci:v2
  - name: nvmrc_only
    ecr: xxxx.dkr.ecr.us-east-1.amazonaws.com/cci:v3
  - name: issue-test
`
	writeErr := os.WriteFile("test-repo-details.yml", []byte(repoYML), 0666)
	if writeErr != nil {
		fmt.Println("Failed to create test file ", writeErr.Error())
		panic("")
	}

	preYML := os.Getenv("MATURITY_REPO_YAML")

	os.Setenv("MATURITY_REPO_YAML", "test-repo-details.yml")
	appConstants = initialize.GetAppConstants()
	repos = getRepos(appConstants.GitHubToken, appConstants.GitHubOwner)
	assert.Len(t, repos, 3)
	assert.Empty(t, repos[2].ECR)
	assert.Equal(t, "xxxx.dkr.ecr.us-east-1.amazonaws.com/ci:v2", repos[0].ECR)
	assert.Equal(t, "xxxx.dkr.ecr.us-east-1.amazonaws.com/cci:v3", repos[1].ECR)
	assert.NotEmpty(t, repos[0].Name)
	assert.Equal(t, "virality", repos[0].Name)
	assert.Equal(t, "issue-test", repos[2].Name)
	assert.Equal(t, "nvmrc_only", repos[1].Name)

	// Revert the old val
	os.Setenv("SCAN_ALL_GITHUB_REPOS", pre)
	os.Setenv("MATURITY_REPO_YAML", preYML)
	os.Remove("test-repo-details.yml")
}
