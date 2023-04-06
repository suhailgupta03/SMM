package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatestPatchDjango_Check(t *testing.T) {
	django := new(LatestPatchDjango)
	app := initialize.GetAppConstants()
	mValue := django.Check(app.Test.Repo.Django)
	assert.Equal(t, types.MaturityValue1, mValue)
}

func TestLatestPatchDjango_Check2(t *testing.T) {
	app := initialize.GetAppConstants()
	g := &github.GitHub{}
	g = g.Init(app.GitHubToken)
	version, found := checkVersionFromRequirementsTxt(g, app.Test.Repo.Django, app.GitHubOwner)
	assert.True(t, found)
	assert.Equal(t, "3.2.15", *version)
}
