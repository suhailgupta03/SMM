package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"fmt"
	"strings"
)

type NodeEOL struct {
}

func extractVersionFromDotXString(dotxNotation string) string {
	split := strings.Split(dotxNotation, ".")
	return split[0]
}

// checkVersionFromEngines checks for "engines" attribute inside package.json and sends
// the mapped nodejs version
func checkVersionFromEngines(packageJson github.PackageJson) (*string, bool) {
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

func checkVersionFromRCFile() (*string, bool) {
	return nil, false
}

func checkVersionFromPackageJson(g *github.GitHub, repoName, gitHubOwner string) (*string, bool) {
	packageJson, err := g.GetPackageJSON(repoName, gitHubOwner)
	if err != nil {
		fmt.Printf("Failed to read package.json for %s\n", repoName)
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

	}

	if existingVersion != nil {
		matchingVersionDetails := util.FindMatchingVersion(*existingVersion, eolDetails)
		if util.IsVersionEOL(*existingVersion, matchingVersionDetails) {
			return types.Yes
		} else {
			return types.No
		}
	}

	return types.NA
}

var Check NodeEOL
