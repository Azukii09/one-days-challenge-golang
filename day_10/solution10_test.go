package day_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {
	assert.Equal(t, 14, ScrabbleScore("cabbage"))
	assert.Equal(t, 10, ScrabbleScore("supper"))
}
