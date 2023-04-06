package main

import (
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type RepoVul struct {
}

func (v *RepoVul) Check(repoName string) types.MaturityCheck {
	repoPath := util.GenerateUrlFromRepoName(repoName)
	isVul, err := util.IsRepoVulnerable(repoPath)
	if err != nil {
		return types.MaturityValue0
	}

	if *isVul {
		return types.MaturityValue1
	} else {
		return types.MaturityValue2
	}
}

func (v *RepoVul) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "No critical vulns: Dependabot",
	}
}

var Check RepoVul
