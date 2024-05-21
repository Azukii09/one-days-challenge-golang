package day_17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NumOfStep(num int) int {
	return num
}

func TestNumOfStep(t *testing.T) {
	assert.Equal(t, 9, NumOfStep(12))
}
