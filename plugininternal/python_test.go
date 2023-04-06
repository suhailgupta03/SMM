package plugininternal

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPythonVersion(t *testing.T) {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	version, found := getVersionFromDockerFile(g, app.Test.Repo.Django, app.GitHubOwner)
	assert.True(t, found)
	assert.Equal(t, "3.10.2-slim", *version)
}
