package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasJSONStrings(t *testing.T) {
	dataList := []string{"apple", "peach", "pear"}
	dl, _ := json.Marshal(dataList)
	dataMap := map[string]string{
		"foo": "bar",
		"bar": "foo",
	}
	dm, _ := json.Marshal(dataMap)

	logStrings := make([]*string, 3)
	str1 := "hello world"
	str2 := "foobar"
	str3 := string(dl)
	logStrings[0] = &str1
	logStrings[1] = &str2
	logStrings[2] = &str3

	hasJSON := hasJSONStrings(logStrings)
	assert.True(t, hasJSON)

	str4 := string(dm)
	logStrings[2] = &str4
	hasJSON = hasJSONStrings(logStrings)
	assert.True(t, hasJSON)

	str5 := "server logs .."
	logStrings[2] = &str5
	hasJSON = hasJSONStrings(logStrings)
	assert.False(t, hasJSON)
}

func TestHasJSONLogging_Check(t *testing.T) {
	app := initialize.GetAppConstants()
	l := new(HasJSONLogging)
	maturity := l.Check(app.Test.AWS.LogGroup + "_" + app.Test.AWS.LogStream)
	assert.NotEqualValues(t, types.MaturityValue0, maturity)

	query := `sort @timestamp desc
| limit 20`

	maturity = l.Check(app.Test.AWS.LogGroup+"_"+app.Test.AWS.LogStream, &query)
	assert.NotEqualValues(t, types.MaturityValue0, maturity)
}

func TestHasJSONLogging_Meta(t *testing.T) {
	l := new(HasJSONLogging)
	meta := l.Meta()
	assert.Equal(t, types.MaturityObservability, meta.Type)
	assert.Equal(t, "Has JSON logging", meta.Name)
}
