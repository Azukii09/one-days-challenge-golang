package day_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotPairNumberForGivenSum(t *testing.T) {
	set := []int{5, 2, 6, 8, 1, 9}
	value := PairNumberForGivenSum(12, set)
	assert.Equal(t, "Pair not found", value)
}
