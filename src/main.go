package main

import (
	"fmt"
)

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity = plantCapacities[0] + plantCapacities[1] +
		plantCapacities[2] + plantCapacities[3] +
		plantCapacities[4] + plantCapacities[5]

	var gridLoad = 75.

	utilization := gridLoad / capacity
	// https://golang.org/pkg/fmt/
	// ""%-20s%.0f"
	fmt.Println("capacity: ", capacity)
	fmt.Println("Load: ", gridLoad)
	fmt.Println("Utilization: ", utilization)
}
