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
			return types.MaturityValue1
		} else {
			return types.MaturityValue2
		}
	}
	return types.MaturityValue0
}

func (r *ReadMe) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDocs,
		Name: "README.markdown",
	}
}

var Check ReadMe
