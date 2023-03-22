package util

import (
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsVersionEOL(t *testing.T) {
	eolDetails, err := http.EOLProvider(http.EOLNode)
	assert.Nil(t, err)
	versionDetails := findMatchingVersion("10", eolDetails)
	assert.NotNil(t, versionDetails)
	assert.Equal(t, "10", versionDetails.Cycle)

	isVersionEOL := isVersionEOL("10", versionDetails)
	assert.True(t, isVersionEOL)

	eolValue := CheckEOL("10", eolDetails)
	assert.Equal(t, types.Yes, eolValue)

	eolValue = CheckEOL("10.2-xyz", eolDetails)
	assert.Equal(t, types.Yes, eolValue)

	/**
	Sample response for django
	  {
	    "cycle": "3.2",
	    "support": "2021-12-01",
	    "eol": "2024-04-01",
	    "latest": "3.2.18",
	    "lts": true,
	    "latestReleaseDate": "2023-02-14",
	    "releaseDate": "2021-04-06"
	  }
	*/
	eolDetails, err = http.EOLProvider(http.EOLDjango)
	assert.Nil(t, err)
	versionDetails = findMatchingVersion("3.2.15", eolDetails)
	assert.NotNil(t, versionDetails)
	assert.Equal(t, "3.2", versionDetails.Cycle)

	eolValue = CheckEOL("3.2.15", eolDetails)
	assert.Equal(t, types.No, eolValue)

	eolDetails, err = http.EOLProvider(http.EOLPython)
	assert.Nil(t, err)
	eolValue = CheckEOL("3.10.2-slim", eolDetails)
	assert.Equal(t, types.No, eolValue)
}

func TestCheckNormalizeString(t *testing.T) {
	version := normalizeVersionString("3.10.2-xyz")
	assert.Equal(t, "3.10.2", version)

	version = normalizeVersionString("1.2.3")
	assert.Equal(t, "1.2.3", version)

	version = normalizeVersionString("slim-1.2.3")
	assert.Equal(t, "1.2.3", version)

	version = normalizeVersionString("1.2.3444")
	assert.Equal(t, "1.2.3444", version)

	version = normalizeVersionString("  1.2  ")
	assert.Equal(t, "1.2", version)
}
