package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"fmt"
)

type LatestPatchDjango struct {
}

func checkVersionFromRequirementsTxt(g *github.GitHub, repoName, gitHubOwner string) (*string, bool) {
	requirements, err := g.GetRequirementsTxt(repoName, gitHubOwner)
	if err != nil {
		fmt.Printf("Warning: Failed to read requirement.txt for %s\n", repoName)
		return nil, false
	}

	version := util.GetVersionFromRequirementsTxt(*requirements, "django")
	if version != nil {
		return version, true
	}
	return nil, false
}

func (lpd LatestPatchDjango) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLDjango)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLDjango + " ")
	}
	var existingVersion *string
	// 1 - Check requirements.txt
	versionFromRequirements, versionFoundFromRequirements := checkVersionFromRequirementsTxt(g, repoName, app.GitHubOwner)
	if versionFoundFromRequirements {
		existingVersion = versionFromRequirements
	}

	if existingVersion != nil {
		mValue := util.IsUsingLatestPatchVersion(*existingVersion, eolDetails)
		return mValue
	}
	return types.MaturityValue0
}

var Check LatestPatchDjango
