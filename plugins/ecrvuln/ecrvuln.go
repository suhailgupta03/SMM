package main

import (
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
)

type ECRVul struct {
}

func (v *ECRVul) Check(ecrImage string) types.MaturityCheck {
	isVuln, err := util.IsImageVulnerable(ecrImage)
	if err != nil {
		return types.MaturityValue0
	}

	if *isVuln {
		return types.MaturityValue1
	} else {
		return types.MaturityValue2
	}
}

func (v *ECRVul) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type:    types.MaturityTypeDependency,
		Name:    "No critical vulns: ECR image",
		EcrType: true,
	}
}

var Check ECRVul
