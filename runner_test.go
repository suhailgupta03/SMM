package main

import (
	"cuddly-eureka-/conf/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigurationLoad(t *testing.T) {
	appConstants = initialize.GetAppConstants()
	assert.NotEmpty(t, appConstants.GitHubOwner, "Tests if the github owner has been loaded")
	assert.NotEmpty(t, appConstants.GitHubToken, "Tests if the github token has been loaded")
}
