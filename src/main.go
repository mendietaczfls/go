package main

import "fmt"

func main() {
	// var plantCapacities []float64
	//
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0, 1}

	gridLoad := 75.
	//
	// var capacity = plantCapacities[0] + plantCapacities[1] +
	// 	plantCapacities[2] + plantCapacities[3] +
	// 	plantCapacities[4] + plantCapacities[5]
	//
	// var gridLoad = 75.
	//
	// utilization := gridLoad / capacity
	// // https://golang.org/pkg/fmt/
	// // ""%-20s%.0f"

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option:")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d Capacity: %.0f\n", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}

		fmt.Println("capacity: ", capacity)
		fmt.Println("Load: ", gridLoad)
		fmt.Println("Utilization: ", gridLoad/capacity*100)

	default:
		fmt.Println("Unknown Option, no action takem")
	}
}
