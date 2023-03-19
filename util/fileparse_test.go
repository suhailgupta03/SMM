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
