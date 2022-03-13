package main

import (
	"BloodPressure/model"
	"fmt"
)

// Test hello
func RunProgram() {
	for i := 0; i <= 1000; i++ {
		model.Create()
	}
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()

	fmt.Println("Done.")
}
