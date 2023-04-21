package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/plugininternal"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type PythonEOL struct {
}

func (python PythonEOL) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLPython)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLPython + " ")
	}

	var existingVersion = plugininternal.FindPythonVersion(repoName)
	if existingVersion != nil {
		eolValue := util.CheckEOL(*existingVersion, eolDetails)
		return eolValue
	}

	return types.MaturityValue0
}

func (python PythonEOL) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Not EOL: Python",
	}
}

var Check PythonEOL
