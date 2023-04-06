package plugininternal

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/util"
	"fmt"
	"regexp"
	"strings"
)

func extractNodeVersionFromNVMRC(nvm string) string {
	/**
	Examples of representation:
		- v14.15.0
		- 14.15.0
	*/
	re := regexp.MustCompile("^v")
	n := re.ReplaceAllString(strings.TrimSpace(nvm), "")
	return n
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
			v := nodejs.(string)
			return &v, true
		}
	}
	return nil, false
}

func checkVersionFromRCFile(g *github.GitHub, repoName, githubOwner string) (*string, bool) {
	nvmrc, err := g.GetDotNVMRC(repoName, githubOwner)
	if err != nil {
		fmt.Printf("Warning: Failed to read .nvmrc for %s\n", repoName)
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

// FindNodeVersion returns the version string as defined in
// package.json or .nvmrc
// example: 3.2.1
func FindNodeVersion(repoName string) *string {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)

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

	return existingVersion
}
