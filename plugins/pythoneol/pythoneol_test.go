package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPythonEOL_Check(t *testing.T) {
	python := new(PythonEOL)
	app := initialize.GetAppConstants()
	mcheck := python.Check(app.Test.Repo.Django)
	assert.Equal(t, types.MaturityValue2, mcheck)
}

func TestPythonEOL_Check2(t *testing.T) {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	version, found := getVersionFromDockerFile(g, app.Test.Repo.Django, app.GitHubOwner)
	assert.True(t, found)
	assert.Equal(t, "3.10.2-slim", *version)
}
