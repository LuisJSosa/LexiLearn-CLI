package test

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("======================================")
	fmt.Println("         Welcome to LexiLearn         ")
	fmt.Println("======================================")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Add a new vocabulary word")
	fmt.Println("2. List all vocabulary words")
	fmt.Println("3. Review vocabulary words")
	fmt.Println("4. Exit")
	fmt.Println("======================================")
	fmt.Print("Enter your choice: ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 4.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		fmt.Println("You chose to add a new vocabulary word.")
		// Future implementation
	case 2:
		fmt.Println("You chose to list all vocabulary words.")
		// Future implementation
	case 3:
		fmt.Println("You chose to review vocabulary words.")
		// Future implementation
	case 4:
		fmt.Println("Exiting... Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		os.Exit(1)
	}
}
