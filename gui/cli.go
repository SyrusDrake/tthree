package gui

import (
	"fmt"
	"tthree/utils"
)

func RunCLI() {
	var celsius float32
	var fahrenheit float32
	fmt.Println("You chose pure CLI. Please enter a number (in °C) to convert to Fahrenheit:")

	_, err := fmt.Scanf("%f", &celsius)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fahrenheit = utils.Calculation(celsius)

	fmt.Printf("%.1f°C is %.1f°F\n", celsius, fahrenheit)
}
