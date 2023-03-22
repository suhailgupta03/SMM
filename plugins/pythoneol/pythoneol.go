package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type PythonEOL struct {
}

func getVersionFromDockerFile(g *github.GitHub, repoName, owner string) (*string, bool) {
	dockerFileContent, dErr := g.GetDockerFile(repoName, owner)
	if dErr != nil {
		return nil, false
	}

	fromCommands := util.ParseDockerFileFromCommand(*dockerFileContent)
	// Filter out the commands where image is python
	var pythonVersion *string = nil
	for _, from := range fromCommands {
		if *from.Image == "python" {
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
func (python PythonEOL) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLPython)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLPython + " ")
	}

	var existingVersion *string
	// 1. Get the content of the dockerfile
	versionFromDF, foundVersionFromDF := getVersionFromDockerFile(g, repoName, app.GitHubOwner)
	if foundVersionFromDF {
		existingVersion = versionFromDF
	}

	if existingVersion != nil {
		eolValue := util.CheckEOL(*existingVersion, eolDetails)
		return eolValue
	}

	return types.NA
}

var Check PythonEOL
