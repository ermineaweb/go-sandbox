package car

import "fmt"

// public interface
type Vehicle interface {
	Accelerate(speed string)
	GetAge() int
}

// private object
type Car struct {
	driverName string
	age        int
}

// public method to make an user value
func MakeVehicle(driverName string, age int) Vehicle {
	return Car{driverName: driverName, age: age}
}

// public method to create a new user
func NewCar(driverName string, age int) *Car {
	return &Car{driverName: driverName, age: age}
}

func (u Car) GetAge() int {
	return u.age
}

func (u Car) Accelerate(speed string) {
	fmt.Printf("%s accelerate to go to %s", u.driverName, speed)
}
