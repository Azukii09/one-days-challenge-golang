package day_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidDNA(t *testing.T) {

	testValue := "GATTACA"
	testResult := DNA{
		A: 3,
		C: 1,
		G: 1,
		T: 2,
	}
	value, err := DNACheckValue(testValue)
	assert.Equal(t, testResult, value)
	assert.Equal(t, "", err)
}

func TestInvalidDNA(t *testing.T) {

	testValue := "INVALID"
	testResult := DNA{
		A: 0,
		C: 0,
		G: 0,
		T: 0,
	}
	value, err := DNACheckValue(testValue)
	assert.Equal(t, testResult, value)
	assert.Equal(t, "error", err)
}
