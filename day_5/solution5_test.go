package day_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanFinish(t *testing.T) {

	car := NewCar(5, 10)
	track := NewTrack(50)

	finishCar := DriveCar(car, track)
	assert.Equal(t, true, finishCar)
}

func TestCantFinish(t *testing.T) {

	car := NewCar(5, 10)
	track := NewTrack(100)

	finishCar := DriveCar(car, track)
	assert.Equal(t, false, finishCar)
}
