package main

import (
	"fmt"
	"tthree/gui"
)

func main() {

	var choice int

	fmt.Println("This is the tthree package main function.")
	fmt.Println("Enter (1) for pure CLI, (2) for BubbleTea, and (3) for Fyne.")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	switch choice {
	case 1:
		gui.RunCLI() // Call the CLI function from the gui package
	case 2:
		gui.RunBubbleTea() // Call the BubbleTea function from the gui package
	case 3:
		gui.RunFyne() // Call the Fyne function from the gui package
	default:
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}
}
