package main

import "cuddly-eureka-/types"

type ECRVul struct {
}

func (v *ECRVul) Check(repoName string) types.MaturityCheck {
	return types.MaturityValue0
}

func (v *ECRVul) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type: types.MaturityTypeDependency,
		Name: "No critical vulns: ECR image",
	}
}

var Check ECRVul
