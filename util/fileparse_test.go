package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersionFromRequirementsTxt(t *testing.T) {
	fileData := `amqp==2.6.0
				billiard==3.6.4.0
				boto==2.49.0
				boto3==1.24.47
				celery==4.4.0
				cffi==1.13.2
				chardet==2.3.0
				codecov==2.0.15
				Django==3.2.15
				coverage==6.4.3`

	version := GetVersionFromRequirementsTxt(fileData, "django")
	assert.Equal(t, "3.2.15", *version)

	version = GetVersionFromRequirementsTxt(fileData, "boto")
	assert.Equal(t, "2.49.0", *version)

	version = GetVersionFromRequirementsTxt(fileData, "foobar")
	assert.Nil(t, version)

	version = GetVersionFromRequirementsTxt("", "django")
	assert.Nil(t, version)
}

func TestGetVersionFromPackageJSON(t *testing.T) {
	pjson := PackageJson{
		"name":    "FooBar",
		"version": "0.0.1",
		"private": true,
		"scripts": map[string]interface{}{
			"start":       "npx react-native start",
			"start:reset": "npx react-native start --reset-cache",
			"test":        "jest",
			"test:watch":  "jest --watch",
			"lint":        "eslint .",
			"android":     "npx react-native run-android",
			"ios":         "npx react-native run-ios",
		},
		"dependencies": map[string]interface{}{
			"react":  "17.0.1",
			"moment": "^2.29.2",
			"axios":  "^0.21.1",
		},
	}

	version := GetVersionFromPackageJSON(pjson, "react")
	assert.NotNil(t, version)
	assert.Equal(t, "17.0.1", *version)

	version = GetVersionFromPackageJSON(pjson, "axios")
	assert.NotNil(t, version)
	assert.Equal(t, "0.21.1", *version)

	version = GetVersionFromPackageJSON(pjson, "foobar")
	assert.Nil(t, version)

	version = GetVersionFromPackageJSON(PackageJson{}, "react")
	assert.Nil(t, version)
}
