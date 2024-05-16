package day_12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanNumber(t *testing.T) {
	assert.Equal(t, "CV", RomanNumbers(105))
	assert.Equal(t, "CVI", RomanNumbers(106))
	assert.Equal(t, "CIV", RomanNumbers(104))
	assert.Equal(t, "MCMXCVI", RomanNumbers(1996))
}
