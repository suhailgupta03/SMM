package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeEOL_Check(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	mcheck := node.Check(app.Test.Repo.Node)
	assert.Equal(t, types.Yes, mcheck, "Should report EOL from package.json")
}

func TestNodeEOL_Check2(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	mcheck := node.Check(app.Test.Repo.NVMRC)
	assert.Equal(t, types.Yes, mcheck, "Should report EOL from .nvmrc")
}

func TestNodeEOL_Check3(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	// tests the response from a repo that neither has package.json nor .nvmrc
	mcheck := node.Check(app.Test.Repo.Empty)
	assert.Equal(t, types.NA, mcheck, "Should report NA type from an empty repo")
}
