package day_17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumOfStep(t *testing.T) {
	assert.Equal(t, 9, NumOfStep(12))
	assert.Equal(t, 4, NumOfStep(16))
	assert.Equal(t, 0, NumOfStep(1))
	assert.Equal(t, 152, NumOfStep(1000000))
}
