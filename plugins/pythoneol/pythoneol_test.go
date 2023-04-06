package main

import (
	"cuddly-eureka-/conf/initialize"
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

func TestPythonEOL_Meta(t *testing.T) {
	python := new(PythonEOL)
	meta := python.Meta()
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
	assert.Equal(t, "Uses latest patch version: Python", meta.Name)
}
