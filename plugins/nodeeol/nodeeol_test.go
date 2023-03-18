package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeEOL_Check(t *testing.T) {
	node := new(NodeEOL)
	app := initialize.GetAppConstants()
	mcheck := node.Check(app.Test.Repo.Node)
	assert.Equal(t, types.Yes, mcheck, "Should report EOL")
}
