package speed

// TODO: define the 'Car' type struct

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	newCar := Car{batteryDrain: batteryDrain, speed: speed, battery:100, distance: 0}
    return newCar
}

// TODO: define the 'Track' type struct

type Track struct {
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{distance: distance}
    return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain {
		car.battery = car.battery - car.batteryDrain
    	car.distance = car.distance + car.speed
    }
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    numberOfDrives := car.battery/car.batteryDrain
	if track.distance <= car.speed*numberOfDrives {
        return true
    }
	return false
}
