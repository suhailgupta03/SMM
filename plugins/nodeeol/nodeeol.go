package main

import "cuddly-eureka-/types"

type NodeEOL struct {
}

func (node NodeEOL) Check(repoPath string) types.MaturityCheck {
	return types.Yes
}

var Check NodeEOL
