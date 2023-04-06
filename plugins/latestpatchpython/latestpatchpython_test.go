package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatestPatchPython_Check(t *testing.T) {
	python := new(LatestPatchPython)
	app := initialize.GetAppConstants()
	mValue := python.Check(app.Test.Repo.Django)
	assert.Equal(t, types.MaturityValue1, mValue)
}

func TestLatestPatchPython_Meta(t *testing.T) {
	python := new(LatestPatchPython)
	meta := python.Meta()
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
	assert.Equal(t, "Uses latest patch version: Python", meta.Name)
}
