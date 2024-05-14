package day_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ScrabbleScore(str string) int {
	return len(str)
}
func TestSuccess(t *testing.T) {
	assert.Equal(t, 14, ScrabbleScore("cabbage"))
}
