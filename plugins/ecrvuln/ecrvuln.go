package main

import "cuddly-eureka-/types"

type ECRVul struct {
}

func (v *ECRVul) Check(repoName string) types.MaturityCheck {
	return types.MaturityValue0
}

var Check ECRVul
