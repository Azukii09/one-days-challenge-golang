package day_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResistorColor(t *testing.T) {
	input1 := []string{"orange", "orange", "black"}
	input2 := []string{"orange", "orange", "orange"}
	assert.Equal(t, "33 ohm", input1)
	assert.Equal(t, "33 kilo ohm", input2)
}
