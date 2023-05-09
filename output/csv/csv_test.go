package csv

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	rowData := make([][]string, 4)
	rowData[0] = []string{"foo", "bar", "foobar"}
	rowData[1] = []string{"alex", "alexa", "alexlexa"}
	rowData[2] = []string{"rose", "flower", "roseflower"}
	rowData[3] = []string{"lotus", "petal", "lotuspetal"}

	rowHeader := []string{"name", "type", "nametype"}

	fileName := "temp.csv"
	generated, err := Generate(rowHeader, rowData, fileName, nil)
	assert.True(t, generated)
	assert.Nil(t, err)

	dir := "some-dir"
	generated, err = Generate(rowHeader, rowData, fileName, &dir)
	assert.True(t, generated)
	assert.Nil(t, err)

	generated, err = Generate(rowHeader, rowData, "some/random/path/"+fileName, nil)
	assert.False(t, generated)
	assert.NotNil(t, err)

	// Remove the file after test is complete
	os.Remove(fileName)
	os.RemoveAll(dir)
}
