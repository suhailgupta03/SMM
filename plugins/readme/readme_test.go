package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadMe_Check(t *testing.T) {
	app := initialize.GetAppConstants()
	readme := new(ReadMe)
	maturityValue := readme.Check(app.Test.Repo.NVMRC) // this repo does not have the readme
	assert.Equal(t, types.MaturityValue2, maturityValue)

	maturityValue = readme.Check(app.Test.Repo.Node) // this repo has the readme
	assert.Equal(t, types.MaturityValue1, maturityValue)

}
