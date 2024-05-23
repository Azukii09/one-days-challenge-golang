package day_19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestISBN(t *testing.T) {
	assert.Equal(t, true, ISBN("3-598-21508-8"))
	assert.Equal(t, false, ISBN("3-598-21508-9"))
	assert.Equal(t, true, ISBN("3-598-21507-X"))
	assert.Equal(t, false, ISBN("3-598-21507-A"))
	assert.Equal(t, false, ISBN("4-598-21507-B"))
	assert.Equal(t, true, ISBN("3598215088"))
}
