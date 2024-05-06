package day_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCS1(t *testing.T) {
	lcs := LCS("ABCBDAB", "BDCABA", 0)
	assert.Equal(t, 4, lcs)
}

func TestLCS2(t *testing.T) {
	lcs := LCS("zfadeg", "cdfsdg", 0)
	assert.Equal(t, 3, lcs)
}
func TestLCS3(t *testing.T) {
	lcs := LCS("abd", "badc", 0)
	assert.Equal(t, 2, lcs)
}
