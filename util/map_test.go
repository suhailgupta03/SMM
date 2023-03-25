package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllKeys(t *testing.T) {
	m := make(map[string]string)
	m["book"] = "Man's eternal quest"
	m["author"] = "Paramahansa Yogananda"
	keys := GetAllKeys(m)
	assert.Len(t, keys, 2)
	assert.Contains(t, keys, "book")
	assert.Contains(t, keys, "author")
}
