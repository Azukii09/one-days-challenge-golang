package day_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaindrop(t *testing.T) {

	assert.Equal(t, "Pling", Raindrop(3))
	assert.Equal(t, "Plang", Raindrop(5))
	assert.Equal(t, "Plong", Raindrop(7))
	assert.Equal(t, "Plong", Raindrop(28))
	assert.Equal(t, "PlingPlang", Raindrop(30))
	assert.Equal(t, "34", Raindrop(34))
}
