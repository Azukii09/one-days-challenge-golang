package day_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipleSum(t *testing.T) {
	finalSum := MultipleSum(20, []int{3, 5})

	assert.Equal(t, 78, finalSum)
}
