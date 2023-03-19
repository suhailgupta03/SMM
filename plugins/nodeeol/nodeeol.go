package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"fmt"
	"regexp"
	"strings"
)

type NodeEOL struct {
}

func extractVersionFromDotXString(dotxNotation string) string {
	split := strings.Split(strings.TrimSpace(dotxNotation), ".")
	return split[0]
}

func extractNodeVersionFromNVMRC(nvm string) string {
	/**
	Examples of representation:
		- v14.15.0
		- 14.15.0
	*/
	re := regexp.MustCompile("^v")
	n := re.ReplaceAllString(strings.TrimSpace(nvm), "")
	return extractVersionFromDotXString(n)
}

// checkVersionFromEngines checks for "engines" attribute inside package.json and sends
// the mapped nodejs version. If the version in the package.json is "14.4.5", the method
// will return "14". Returns false, if the version was not found
func checkVersionFromEngines(packageJson util.PackageJson) (*string, bool) {
	/**
	Example:
	  "engines": {
	    "node": "10.x"
	  }
	*/
	engines, found := packageJson["engines"]
	if found {
		nodejs, fNode := engines.(map[string]interface{})["node"]
		if fNode {
			existingVersion := extractVersionFromDotXString(nodejs.(string))
			return &existingVersion, true
		}
	}
	return nil, false
}

func checkVersionFromRCFile(g *github.GitHub, repoName, githubOwner string) (*string, bool) {
	nvmrc, err := g.GetDotNVMRC(repoName, githubOwner)
	if err != nil {
		fmt.Printf("Failed to read .nvmrc for %s\n", repoName)
		return nil, false
	}

	version := extractNodeVersionFromNVMRC(*nvmrc)
	if version != "" {
		return &version, true
	}
	return nil, false
}

func checkVersionFromPackageJson(g *github.GitHub, repoName, gitHubOwner string) (*string, bool) {
	packageJson, err := g.GetPackageJSON(repoName, gitHubOwner)
	if err != nil {
		fmt.Printf("Warning: Failed to read package.json for %s\n", repoName)
		return nil, false
	}

	versionFromEngines, foundFromEngine := checkVersionFromEngines(packageJson)
	if foundFromEngine {
		return versionFromEngines, true
	}
	return nil, false
}

func (node NodeEOL) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLNode)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLNode + " ")
	}
	var existingVersion *string

	// 1 - Check package.json
	versionFromPJ, foundFromPJ := checkVersionFromPackageJson(g, repoName, app.GitHubOwner)
	if foundFromPJ {
		existingVersion = versionFromPJ
	}

	// 2 - Check .nvmrc
	if existingVersion == nil {
		versionFromNVMRC, foundFromNVMRC := checkVersionFromRCFile(g, repoName, app.GitHubOwner)
		if foundFromNVMRC {
			existingVersion = versionFromNVMRC
		}
	}

	if existingVersion != nil {
		eolValue := util.CheckEOL(*existingVersion, eolDetails)
		return eolValue
	}

	return types.NA
}

var Check NodeEOL
