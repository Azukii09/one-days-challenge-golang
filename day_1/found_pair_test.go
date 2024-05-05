package day_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairNumberForGivenSum(t *testing.T) {
	set := []int{8, 7, 2, 5, 3, 1}
	value := PairNumberForGivenSum(10, set)
	assert.Equal(t, "Pair found 8,2 or Pair found 7,3", value)
}
