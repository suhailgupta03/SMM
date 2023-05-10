package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/plugininternal"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type LatestPatchNode struct {
}

func (lpn LatestPatchNode) Check(repoName string, opts ...*string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLNode)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLNode + " ")
	}
	var existingVersion = plugininternal.FindNodeVersion(repoName)
	if existingVersion != nil {
		return util.IsUsingLatestPatchVersion(*existingVersion, eolDetails)
	}

	return types.MaturityValue0
}

func (lpn LatestPatchNode) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Uses latest patch version: Node",
	}
}

var Check LatestPatchNode
