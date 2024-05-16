package day_12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanNumber(t *testing.T) {
	numbers1, err1 := RomanNumbers(105)
	if err1 != nil {
		return
	}
	numbers2, err2 := RomanNumbers(106)
	if err2 != nil {
		return
	}
	numbers3, err3 := RomanNumbers(104)
	if err3 != nil {
		return
	}
	numbers4, err4 := RomanNumbers(1996)
	if err4 != nil {
		return
	}
	assert.Equal(t, "CV", numbers1)
	assert.Equal(t, "CVI", numbers2)
	assert.Equal(t, "CIV", numbers3)
	assert.Equal(t, "MCMXCVI", numbers4)
}
