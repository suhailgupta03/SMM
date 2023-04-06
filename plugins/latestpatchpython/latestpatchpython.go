package main

import (
	"cuddly-eureka-/http"
	"cuddly-eureka-/plugininternal"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type LatestPatchPython struct {
}

func (lpp LatestPatchPython) Check(repoName string) types.MaturityCheck {
	eolDetails, eolErr := http.EOLProvider(http.EOLPython)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLPython + " ")
	}

	var existingVersion = plugininternal.FindPythonVersion(repoName)
	if existingVersion != nil {
		return util.IsUsingLatestPatchVersion(*existingVersion, eolDetails)
	}
	return types.MaturityValue0
}

func (lpp LatestPatchPython) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Uses latest patch version: Python",
	}
}

var Check LatestPatchPython
