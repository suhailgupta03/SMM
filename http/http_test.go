package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEOLProvider(t *testing.T) {
	prodDetails, err := EOLProvider(EOLNode)
	assert.Nil(t, err, "EOL provider must not return an error")
	assert.Greater(t, len(prodDetails), 1)
}
