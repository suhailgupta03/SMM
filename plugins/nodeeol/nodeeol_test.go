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
	assert.Equal(t, types.MaturityValue1, mcheck, "Should report EOL from package.json")
}

func TestNodeEOL_Check2(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	mcheck := node.Check(app.Test.Repo.NVMRC)
	assert.Equal(t, types.MaturityValue1, mcheck, "Should report EOL from .nvmrc")
}

func TestNodeEOL_Check3(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	// tests the response from a repo that neither has package.json nor .nvmrc
	mcheck := node.Check(app.Test.Repo.Empty)
	assert.Equal(t, types.MaturityValue0, mcheck, "Should report MaturityValue0 type from an empty repo")
}

func TestExtractVersionFromDotString(t *testing.T) {
	cycle := extractCycleFromDotXString("13.3.1")
	assert.Equal(t, "13", cycle)

	cycle = extractCycleFromDotXString("")
	assert.Equal(t, "", cycle)

	cycle = extractCycleFromDotXString("  ")
	assert.Equal(t, "", cycle)
}

func TestNodeEOL_Meta(t *testing.T) {
	node := new(NodeEOL)
	meta := node.Meta()
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
	assert.Equal(t, "Uses latest patch version: Node", meta.Name)
}
