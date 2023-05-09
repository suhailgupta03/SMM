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
	assert.Equal(t, types.MaturityValue1, eolValue)

	eolValue = CheckEOL("10.2-xyz", eolDetails)
	assert.Equal(t, types.MaturityValue1, eolValue)

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

	versionDetails = findMatchingVersion("1.9.5", eolDetails)
	assert.NotNil(t, versionDetails)
	assert.Equal(t, "-1", versionDetails.Cycle)

	eolValue = CheckEOL("3.2.15", eolDetails)
	assert.Equal(t, types.MaturityValue2, eolValue)

	eolDetails, err = http.EOLProvider(http.EOLPython)
	assert.Nil(t, err)
	eolValue = CheckEOL("3.10.2-slim", eolDetails)
	assert.Equal(t, types.MaturityValue2, eolValue)

	eolDetails, err = http.EOLProvider(http.EOLGO)
	assert.Nil(t, err)
	eolValue = CheckEOL("1.10", eolDetails)
	assert.Equal(t, types.MaturityValue1, eolValue)
	eolValue = CheckEOL("1.20", eolDetails)
	assert.Equal(t, types.MaturityValue2, eolValue)

	eolDetails, err = http.EOLProvider(http.EOLReact)
	assert.Nil(t, err)
	eolValue = CheckEOL("17.0.1", eolDetails)
	assert.Equal(t, types.MaturityValue2, eolValue)
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

	version = normalizeVersionString("1.2.3.")
	assert.Equal(t, "1.2.3", version)
}

func TestIsUsingLatestPatchVersion(t *testing.T) {
	eolDetails, err := http.EOLProvider(http.EOLDjango)
	assert.Nil(t, err)
	mValue := IsUsingLatestPatchVersion("1.11.29", eolDetails)
	assert.Equal(t, types.MaturityValue2, mValue)
	mValue = IsUsingLatestPatchVersion("1.11.28", eolDetails)
	assert.Equal(t, types.MaturityValue1, mValue)
	mValue = IsUsingLatestPatchVersion("", eolDetails)
	assert.Equal(t, types.MaturityValue0, mValue)
	eolDetails, err = http.EOLProvider(http.EOLPython)
	assert.Nil(t, err)
	mValue = IsUsingLatestPatchVersion("3.10.11", eolDetails)
	assert.Equal(t, types.MaturityValue2, mValue)
	mValue = IsUsingLatestPatchVersion("3.10.10", eolDetails)
	assert.Equal(t, types.MaturityValue1, mValue)
}
