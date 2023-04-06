package plugininternal

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/util"
)

func getVersionFromDockerFile(g *github.GitHub, repoName, owner string) (*string, bool) {
	dockerFileContent, dErr := g.GetDockerFile(repoName, owner)
	if dErr != nil {
		return nil, false
	}

	fromCommands := util.ParseDockerFileFromCommand(*dockerFileContent)
	// Filter out the commands where image is python
	var pythonVersion *string = nil
	for _, from := range fromCommands {
		if from.Image != nil && *from.Image == "python" {
			pythonVersion = from.Tag
			// Will break as soon as we find the first version of python
			break
		}
	}

	if pythonVersion != nil {
		return pythonVersion, true
	}

	return pythonVersion, false
}

func FindPythonVersion(repoName string) *string {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)

	var existingVersion *string

	// 1. Get the content of the dockerfile
	versionFromDF, foundVersionFromDF := getVersionFromDockerFile(g, repoName, app.GitHubOwner)
	if foundVersionFromDF {
		existingVersion = versionFromDF
	}

	return existingVersion
}
