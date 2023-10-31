package main

import (
	"fmt"
	"strings"
)

func main() {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	var number1, number2 int
	var operation string

	for {
		fmt.Println("Enter the first number:")
		if _, err := fmt.Scanln(&number1); err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		fmt.Println("Enter the operation (+, -, *, /):")
		if _, err := fmt.Scanln(&operation); err != nil {
			fmt.Println("Invalid input. Please enter a valid operat	ion.")
			continue
		}

		operation = strings.TrimSpace(operation)

		if _, ok := ops[operation]; !ok {
			fmt.Println("Invalid operation. Please enter a valid operation (+, -, *, /).")
			continue
		}

		fmt.Println("Enter the second number:")
		if _, err := fmt.Scanln(&number2); err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		break
	}

	result := ops[operation](number1, number2)
	fmt.Printf("Result: %d %s %d = %d\n", number1, operation, number2, result)
}
