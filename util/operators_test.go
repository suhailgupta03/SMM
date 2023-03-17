package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOR(t *testing.T) {
	result := GetOR("val1", "val2")
	assert.Equal(t, "val1", result)

	result = GetOR("", "")
	assert.Equal(t, "", result)

	result = GetOR("", "val2")
	assert.Equal(t, "val2", result)

	result = GetOR("val1", "")
	assert.Equal(t, "val1", result)
}
