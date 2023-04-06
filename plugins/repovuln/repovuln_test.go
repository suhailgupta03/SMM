package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepoVul_Check(t *testing.T) {
	vplugin := new(RepoVul)
	app := initialize.GetAppConstants()
	mval := vplugin.Check(app.Test.Repo.Trivy)
	assert.Equal(t, types.MaturityValue1, mval)

	mval = vplugin.Check("ISVqHDvVBRLniKZOxRxN")
	assert.Equal(t, types.MaturityValue0, mval)
}

func TestRepoVul_Meta(t *testing.T) {
	v := new(RepoVul)
	meta := v.Meta()
	assert.Equal(t, "No critical vulns: Dependabot", meta.Name)
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
}
