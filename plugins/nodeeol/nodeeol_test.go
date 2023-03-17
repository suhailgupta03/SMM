package main

import (
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	PackageJsonRepoWithEngines = "issue-test"
)

func TestNodeEOL_Check(t *testing.T) {
	node := new(NodeEOL)
	mcheck := node.Check(PackageJsonRepoWithEngines)
	assert.Equal(t, types.Yes, mcheck, "Should report EOL")
}
