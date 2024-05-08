package day_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationCheckSuccess(t *testing.T) {
	assert.Equal(t, true, PermutationCheck("abc", "cba"))
	assert.Equal(t, true, PermutationCheck("party", "tryap"))
}
func TestPermutationCheckFailed(t *testing.T) {
	assert.Equal(t, false, PermutationCheck("abcd", "cba"))
	assert.Equal(t, false, PermutationCheck("parzy", "tryap"))
}
