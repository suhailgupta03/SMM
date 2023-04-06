package main

import (
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestECRVul_Meta(b *testing.T) {
	ecr := new(ECRVul)
	meta := ecr.Meta()
	assert.Equal(b, types.MaturityTypeDependency, meta.Type)
	assert.Equal(b, "No critical vulns: ECR image", meta.Name)
}
