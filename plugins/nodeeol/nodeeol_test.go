package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"cuddly-eureka-/util"
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

func TestNodeEOL_Check4(t *testing.T) {
	version := extractVersionFromDotXString("13.3.1")
	assert.Equal(t, "13", version)

	version = extractVersionFromDotXString("")
	assert.Equal(t, "", version)

	version = extractVersionFromDotXString("  ")
	assert.Equal(t, "", version)
}

func TestNodeEOL_Check5(t *testing.T) {
	version := extractNodeVersionFromNVMRC("v14.5.1")
	assert.Equal(t, "14", version)

	version = extractNodeVersionFromNVMRC("")
	assert.Equal(t, "", version)

	version = extractNodeVersionFromNVMRC(" ")
	assert.Equal(t, "", version)
}

func TestNodeEOL_Check6(t *testing.T) {
	packageJson := util.PackageJson{
		"name":        "foo-bar",
		"version":     "1.0.0",
		"description": "",
		"main":        "index.js",
		"engines": map[string]interface{}{
			"node": "15.4.5",
		},
	}
	version, found := checkVersionFromEngines(packageJson)
	assert.True(t, found, "Should be able to find the nodejs version from package.json")
	assert.Equal(t, "15", *version)

	packageJson = util.PackageJson{
		"name":        "foo-bar",
		"version":     "1.0.0",
		"description": "",
		"main":        "index.js",
	}
	version, found = checkVersionFromEngines(packageJson)
	assert.False(t, found)
	assert.Nil(t, version)
}
