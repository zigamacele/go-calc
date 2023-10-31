package utils

import (
	"fmt"
	"strings"
)

func ReadInt() int {
	var number int
	if _, err := fmt.Scanln(&number); err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return ReadInt()
	}
	return number
}

func ReadOperation(ops map[string]func(int, int) int) string {
	var operation string
	if _, err := fmt.Scanln(&operation); err != nil || ops[operation] == nil {
		fmt.Println("Invalid input. Please enter a valid operation (+, -, *, /).")
		return ReadOperation(ops)
	}
	operation = strings.TrimSpace(operation)
	return operation
}

func ValidateNumber(number int) bool {
	return number >= 0
}
