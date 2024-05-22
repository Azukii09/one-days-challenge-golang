package day_18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Valid(num string) bool {
	return false
}

func TestValid(t *testing.T) {
	assert.Equal(t, false, Valid("1"))
	assert.Equal(t, false, Valid("0"))
	assert.Equal(t, true, Valid("059"))
	assert.Equal(t, true, Valid("59"))
	assert.Equal(t, true, Valid("055 444 285"))
	assert.Equal(t, false, Valid("055 444 286"))
}
