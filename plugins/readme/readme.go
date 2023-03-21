package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/types"
	"fmt"
)

type ReadMe struct {
}

func (r *ReadMe) Check(repoName string) types.MaturityCheck {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)

	exists, err := g.DoesReadMeExist(repoName, app.GitHubOwner)
	if err != nil {
		fmt.Printf("Warning: Failed to check for the existence of readme for repo %s\n", repoName)
	} else {
		if *exists {
			return types.Yes
		} else {
			return types.No
		}
	}
	return types.NA
}

var Check ReadMe
