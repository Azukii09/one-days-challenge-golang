package day_5

type RemoteCar struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) RemoteCar {
	return RemoteCar{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func DriveCar(car RemoteCar, track Track) bool {
	for i := 100; i >= 0; i -= car.batteryDrain {
		car.distance += car.speed
	}

	if car.distance >= track.distance {
		return true
	}
	return false
}
