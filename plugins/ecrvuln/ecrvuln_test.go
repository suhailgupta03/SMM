package main

import (
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestECRVul_Check(t *testing.T) {
	ecr := new(ECRVul)
	mValue := ecr.Check("random-image-name")
	assert.Equal(t, types.MaturityValue0, mValue)

	mValue = ecr.Check("python:3.4") // testing with a non-ecr image
	assert.Equal(t, types.MaturityValue1, mValue)

}

func TestECRVul_Meta(b *testing.T) {
	ecr := new(ECRVul)
	meta := ecr.Meta()
	assert.Equal(b, types.MaturityTypeDependency, meta.Type)
	assert.Equal(b, "No critical vulns: ECR image", meta.Name)
	assert.True(b, meta.EcrType)
}
