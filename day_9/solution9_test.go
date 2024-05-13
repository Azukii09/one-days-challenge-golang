package day_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResistorColor(t *testing.T) {
	input1 := []string{"orange", "orange", "black"}
	input2 := []string{"orange", "orange", "brown"}
	input3 := []string{"orange", "orange", "red"}
	input4 := []string{"orange", "orange", "orange"}
	input5 := []string{"orange", "orange", "yellow"}
	input6 := []string{"orange", "orange", "green"}
	input7 := []string{"orange", "orange", "blue"}
	input8 := []string{"orange", "orange", "violet"}
	input9 := []string{"orange", "orange", "grey"}
	input10 := []string{"orange", "orange", "white"}

	assert.Equal(t, "33 ohm", ResistorColorConverter(input1))
	assert.Equal(t, "330 ohm", ResistorColorConverter(input2))
	assert.Equal(t, "3300 ohm", ResistorColorConverter(input3))
	assert.Equal(t, "33 kilo ohm", ResistorColorConverter(input4))
	assert.Equal(t, "330 kilo ohm", ResistorColorConverter(input5))
	assert.Equal(t, "3300 kilo ohm", ResistorColorConverter(input6))
	assert.Equal(t, "33 mega ohm", ResistorColorConverter(input7))
	assert.Equal(t, "330 mega ohm", ResistorColorConverter(input8))
	assert.Equal(t, "3300 mega ohm", ResistorColorConverter(input9))
	assert.Equal(t, "33 giga ohm", ResistorColorConverter(input10))
}
