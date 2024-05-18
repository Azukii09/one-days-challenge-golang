package day_14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArmstrongNumber(t *testing.T) {
	assert.Equal(t, true, IsNumber(9))
	assert.Equal(t, true, IsNumber(153))
	assert.Equal(t, false, IsNumber(10))
	assert.Equal(t, false, IsNumber(154))
}
