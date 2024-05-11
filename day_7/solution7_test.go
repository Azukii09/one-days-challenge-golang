package day_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {

	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Alan"},
		Developer{Name: "Jennifer"},
		Developer{Name: "Graham"},
		Developer{Name: "Paul"},
		Developer{Name: "Alan"},
	}
	outputTest := FilterUnique(input)
	resultTest := []string{
		"Elliot",
		"Alan",
		"Jennifer",
		"Graham",
		"Paul",
	}
	assert.Equal(t, resultTest, outputTest)
}
