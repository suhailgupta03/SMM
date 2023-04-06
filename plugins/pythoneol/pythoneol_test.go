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
