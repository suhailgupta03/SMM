package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCov_IsCodeCovActivated(t *testing.T) {
	var test = initialize.GetAppConstants().Test
	activated, err := isCodeCovActivated(test.CodeCov.RepoName, test.CodeCov.RepoOwner, "github", test.CodeCov.BearerToken)
	assert.NotNil(t, activated)
	assert.Nil(t, err)
	assert.True(t, *activated)

	activated, err = isCodeCovActivated(test.CodeCov.RepoName, test.CodeCov.RepoOwner, "github", "..")
	assert.Nil(t, activated)
	assert.NotNil(t, err)
}

func TestMinCov_Check(t *testing.T) {
	var test = initialize.GetAppConstants().Test
	c := new(MinCov)
	input := fmt.Sprintf("%s_%s_%s_%s", test.CodeCov.RepoName, test.CodeCov.RepoOwner, "github", test.CodeCov.BearerToken)
	maturity := c.Check(input)
	assert.Equal(t, types.MaturityValue2, maturity)

	maturity = c.Check("a_b_c_d")
	assert.Equal(t, types.MaturityValue0, maturity)

	maturity = c.Check("a_b")
	assert.Equal(t, types.MaturityValue0, maturity)
}

func TestMinCov_Meta(t *testing.T) {
	c := new(MinCov)
	m := c.Meta()
	assert.Equal(t, types.MaturityCI, m.Type)
	assert.True(t, m.CodeCovType)
	assert.Equal(t, "Unit tests with min coverage enforced", m.Name)
}
