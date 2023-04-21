package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasLogging_Check(t *testing.T) {
	app := initialize.GetAppConstants()
	l := new(HasLogging)
	maturity := l.Check(app.Test.AWS.LogGroup + "_" + app.Test.AWS.LogStream)
	assert.Equal(t, types.MaturityValue2, maturity)
}

func TestHasLogging_Meta(t *testing.T) {
	l := new(HasLogging)
	meta := l.Meta()
	assert.Equal(t, types.MaturityObservability, meta.Type)
	assert.Equal(t, "Has (some) logging", meta.Name)
}
