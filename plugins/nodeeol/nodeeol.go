package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	nodeinternal "cuddly-eureka-/plugininternal"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"strings"
)

type NodeEOL struct {
}

func extractCycleFromDotXString(dotxNotation string) string {
	split := strings.Split(strings.TrimSpace(dotxNotation), ".")
	return split[0]
}

func (node NodeEOL) Check(repoName string, opts ...*string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLNode)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLNode + " ")
	}
	var existingVersion = nodeinternal.FindNodeVersion(repoName)
	if existingVersion != nil {
		cycleNumber := extractCycleFromDotXString(*existingVersion)
		eolValue := util.CheckEOL(cycleNumber, eolDetails)
		return eolValue
	}

	return types.MaturityValue0
}

func (node NodeEOL) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Not EOL: Node",
	}
}

var Check NodeEOL
