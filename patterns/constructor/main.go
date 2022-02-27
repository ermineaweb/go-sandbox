package main

import (
	"fmt"
	"go-sandbox/patterns/constructor/car"
)

func main() {
	harryCar := car.NewCar("harry", 3)
	harryCar.Accelerate("140km/h")
	fmt.Printf("harry's car is %d years old", harryCar.GetAge())
}
