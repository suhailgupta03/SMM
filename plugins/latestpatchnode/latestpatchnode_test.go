package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatestPatchNode_Check(t *testing.T) {
	node := new(LatestPatchNode)
	app := initialize.GetAppConstants()
	mValue := node.Check(app.Test.Repo.Node)
	assert.Equal(t, types.MaturityValue1, mValue)
}

func TestLatestPatchNode_Check2(t *testing.T) {
	node := new(LatestPatchNode)
	app := initialize.GetAppConstants()
	mValue := node.Check(app.Test.Repo.NVMRC)
	assert.Equal(t, types.MaturityValue2, mValue)
}
