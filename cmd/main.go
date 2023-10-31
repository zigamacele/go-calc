package main

import (
	"calc/internal/utils"
	"fmt"
)

func main() {
	ops := utils.GetOperations()

	var number1, number2 int
	var operation string

	for {
		fmt.Println("Enter the first number:")
		number1 = utils.ReadInt()

		fmt.Println("Enter the operation (+, -, *, /):")
		operation = utils.ReadOperation(ops)

		fmt.Println("Enter the second number:")
		number2 = utils.ReadInt()

		if utils.ValidateNumber(number1) && utils.ValidateNumber(number2) {
			break
		}
	}

	result := ops[operation](number1, number2)
	fmt.Printf("Result: %d %s %d = %d\n", number1, operation, number2, result)
}
