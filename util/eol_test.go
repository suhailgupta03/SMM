package util

import (
	"cuddly-eureka-/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsVersionEOL(t *testing.T) {
	eolDetails, err := http.EOLProvider(http.EOLNode)
	assert.Nil(t, err)
	versionDetails := FindMatchingVersion("10", eolDetails)
	assert.NotNil(t, versionDetails)
	assert.Equal(t, "10", versionDetails.Cycle)
	isVersionEOL := IsVersionEOL("10", versionDetails)
	assert.True(t, isVersionEOL)
}
