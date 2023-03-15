package main

import (
	"cuddly-eureka-/types"
	"fmt"
)

type ExtendedMaturityCheck types.MaturityCheck

func (pv *ExtendedMaturityCheck) Print(repoName string, checkName string, checkValue types.MaturityCheck) {
	fmt.Printf(checkName+" in "+repoName+" has a maturity value of %d \n", checkValue)
}
