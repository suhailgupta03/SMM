package util

import (
	"cuddly-eureka-/conf/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUrlFromRepoName(t *testing.T) {
	app := initialize.GetAppConstants()
	repoPath := GenerateUrlFromRepoName("my-repo")
	assert.Equal(t, "https://github.com/"+app.GitHubOwner+"/my-repo", repoPath)
}
