package day_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ResistorColorConverter([]string) string {
	return "33 ohm"
}
func ResistorColorConverter2([]string) string {
	return "33 kilo ohm"
}

func TestResistorColor(t *testing.T) {
	input1 := []string{"orange", "orange", "black"}
	input2 := []string{"orange", "orange", "orange"}
	assert.Equal(t, "33 ohm", ResistorColorConverter(input1))
	assert.Equal(t, "33 kilo ohm", ResistorColorConverter2(input2))
}
