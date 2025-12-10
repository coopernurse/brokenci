package main

import (
	"fmt"
	"os"

	"github.com/coopernurse/brokenci/internal/calculator"
	"github.com/coopernurse/brokenci/internal/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: pretest <operation> <num1> <num2>")
		fmt.Println("Operations: add, subtract, multiply, divide")
		os.Exit(1)
	}

	operation := os.Args[1]
	a, b := utils.ParseNumbers(os.Args[2], os.Args[3])

	var result float64
	var err error

	switch operation {
	case "add":
		result = calculator.Add(a, b)
	case "subtract":
		result = calculator.Subtract(a, b)
	case "multiply":
		result = calculator.Multiply(a, b)
	case "divide":
		result, err = calculator.Divide(a, b)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown operation: %s\n", operation)
		os.Exit(1)
	}

	fmt.Printf("This is the result: %v\n", result)
}
