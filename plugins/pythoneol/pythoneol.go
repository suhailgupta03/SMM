package main

import "cuddly-eureka-/types"

type PythonEOL struct {
}

func (python PythonEOL) Check(repoPath string) types.MaturityCheck {
	return types.No
}

var Check PythonEOL
