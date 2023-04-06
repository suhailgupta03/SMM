package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"fmt"
)

type ReactEOL struct {
}

func getVersionFromPackageJson(g *github.GitHub, repoName, gitHubOwner string) (*string, bool) {
	packageJson, err := g.GetPackageJSON(repoName, gitHubOwner)
	if err != nil {
		fmt.Printf("Warning: Failed to read package.json for %s\n", repoName)
		return nil, false
	}

	version := util.GetVersionFromPackageJSON(packageJson, "react")
	if version != nil {
		return version, true
	}

	return nil, false
}
func (react *ReactEOL) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLReact)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLReact + " ")
	}

	var existingVersion *string

	// 1 - Check package.json
	versionFromPJ, foundVersionFromPJ := getVersionFromPackageJson(g, repoName, app.GitHubOwner)
	if foundVersionFromPJ {
		existingVersion = versionFromPJ
	}

	if existingVersion != nil {
		eolValue := util.CheckEOL(*existingVersion, eolDetails)
		return eolValue
	}

	return types.MaturityValue0
}

func (react *ReactEOL) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Uses latest patch version: React",
	}
}

var Check ReactEOL
