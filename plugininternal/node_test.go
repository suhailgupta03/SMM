package plugininternal

import (
	"cuddly-eureka-/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractVersionFromNVMRC(t *testing.T) {
	version := extractNodeVersionFromNVMRC("v14.5.1")
	assert.Equal(t, "14.5.1", version)

	version = extractNodeVersionFromNVMRC("")
	assert.Equal(t, "", version)

	version = extractNodeVersionFromNVMRC(" ")
	assert.Equal(t, "", version)
}

func TestCheckVersionFromEngines(t *testing.T) {
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
	assert.Equal(t, "15.4.5", *version)

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
