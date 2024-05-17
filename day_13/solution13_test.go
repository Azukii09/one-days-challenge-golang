package day_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipleSum(t *testing.T) {
	finalSum := MultipleSum(20, []int{3, 5})
	finalSum2 := MultipleSum(15, []int{2, 3, 4})
	assert.Equal(t, 78, finalSum)
	assert.Equal(t, 68, finalSum2)
}
