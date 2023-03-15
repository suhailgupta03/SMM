package main

import "cuddly-eureka-/types"

type DjangoEOL struct {
}

func (django DjangoEOL) Check(repoPath string) types.MaturityCheck {
	return types.NA
}

var Check DjangoEOL
