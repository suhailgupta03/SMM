package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
	"fmt"
	"regexp"
)

type GoEOL struct {
}

func extractVersionFromMod(file *string) *string {
	var version *string
	re := regexp.MustCompile(`(go)\s+(\d+\.\d+)`)
	if file != nil && re.MatchString(*file) {
		groups := re.FindStringSubmatch(*file)
		if len(groups) == 3 {
			version = &groups[2]
		}
	}

	return version
}

func checkVersionFromModFile(g *github.GitHub, repoName, owner string) (*string, bool) {
	mod, err := g.GetGoModFile(repoName, owner)
	if err != nil {
		fmt.Printf("Warning: Failed to read go.mod for %s\n", repoName)
		return nil, false
	}

	version := extractVersionFromMod(mod)
	if version != nil {
		return version, true
	}

	return nil, false
}

func (g *GoEOL) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	git := &github.GitHub{}
	git = git.Init(app.GitHubToken)
	eolDetails, eolErr := http.EOLProvider(http.EOLGO)
	if eolErr != nil {
		panic("Failed to find EOL details for " + http.EOLDjango + " ")
	}
	var existingVersion *string
	// 1. Check version from mod file
	versionFromMod, versionFoundFromMod := checkVersionFromModFile(git, repoName, app.GitHubOwner)
	if versionFoundFromMod {
		existingVersion = versionFromMod
	}

	if existingVersion != nil {
		eolValue := util.CheckEOL(*existingVersion, eolDetails)
		return eolValue
	}
	return types.MaturityValue0
}

func (g *GoEOL) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "Not EOL: GO",
	}
}
