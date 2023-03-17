package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsDateGreater(t *testing.T) {
	firstDate := time.Now()
	h, _ := time.ParseDuration("1h")
	secondDate := firstDate.Add(h)
	isGreater := IsDateGreater(firstDate, secondDate)
	assert.False(t, isGreater)
	ph, _ := time.ParseDuration("-1h")
	secondDate = firstDate.Add(ph)
	isGreater = IsDateGreater(firstDate, secondDate)
	assert.True(t, isGreater)
}
