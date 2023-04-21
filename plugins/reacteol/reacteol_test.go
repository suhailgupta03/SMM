package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReactEOL_Check(t *testing.T) {
	react := new(ReactEOL)
	app := initialize.GetAppConstants()
	mcheck := react.Check(app.Test.Repo.Node)
	assert.Equal(t, types.MaturityValue2, mcheck, "Should not report EOL from package.json")
}

func TestReactEOL_Check2(t *testing.T) {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	version, found := getVersionFromPackageJson(g, app.Test.Repo.Node, app.GitHubOwner)
	assert.True(t, found)
	assert.Equal(t, "18.2.0", *version)
}

func TestReactEOL_Meta(t *testing.T) {
	react := new(ReactEOL)
	meta := react.Meta()
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
	assert.Equal(t, "Not EOL: React", meta.Name)
}
