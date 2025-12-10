package utils

import (
	"fmt"
)

// Lint Issue 1: Long function (should be refactored)
func VeryLongFunctionThatShouldBeRefactored() {
	fmt.Println("Starting...")
	var x int
	x = 1
	x = x + 1
	x = x * 2
	fmt.Println(x)
	x = x - 1
	fmt.Println(x)
	x = x / 2
	fmt.Println(x)
	x = x % 3
	fmt.Println(x)
	x = x << 1
	fmt.Println(x)
	x = x >> 1
	fmt.Println(x)
	x = x | 1
	fmt.Println(x)
	x = x & 1
	fmt.Println(x)
	x = x ^ 1
	fmt.Println(x)
	// ... and 20 more lines of trivial operations
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Still going...")
	// Another 15 lines
	y := 0
	for j := 0; j < 5; j++ {
		y += j
		fmt.Println(y)
	}
	fmt.Println("Almost done...")
	// Final 10 lines
	z := "end"
	fmt.Println(z)
}

// Lint Issue 2: Cognitive complexity (nested loops and conditions)
func HighCognitiveComplexity(condition bool, items []string) {
	if condition {
		fmt.Println("Condition is true")
		for i := 0; i < len(items); i++ {
			if items[i] != "" {
				for j := 0; j < i; j++ {
					if j%2 == 0 {
						switch items[j] {
						case "a":
							fmt.Println("Found a")
						case "b":
							fmt.Println("Found b")
							if i > 5 {
								fmt.Println("i > 5")
							}
						default:
							fmt.Println("Default")
						}
					}
				}
			}
		}
	} else {
		fmt.Println("Condition is false")
	}
}

// Lint Issue 3: Magic numbers
func CalculatePrice(quantity int) float64 {
	tax := 0.0825      // Magic number: tax rate
	shipping := 5.99   // Magic number: shipping cost
	discount := 0.15   // Magic number: discount rate
	basePrice := 19.99 // Magic number: base price

	total := (float64(quantity)*basePrice)*(1.0-discount) + shipping
	return total * (1.0 + tax)
}

// Lint Issue 4: Unnecessary leading newline in error message
func ValidateInput(input string) error {
	if input == "" {
		return fmt.Errorf(`
Input cannot be empty`) // Bad: leading newline in error
	}
	if len(input) > 100 {
		return fmt.Errorf("input is too long (max 100 chars)")

	}
	return nil
}

// Lint Issue 5: Inconsistent naming (should be camelCase for const)
const MAX_RETRIES = 3 // Should be MaxRetries

// Lint Issue 6: Function with too many parameters
func TooManyParameters(a, b, c, d, e, f, g, h int, i, j, k string) int {
	return a + b + c + d + e + f + g + h
}

// Lint Issue 7: Missing period in comment
// CalculateSum adds two numbers
// This function is used throughout the app  // Missing period at end
func CalculateSum(a, b int) int {
	return a + b
}

// Lint Issue 8: Ineffective break statement
func FindFirstEven(numbers []int) int {
	for _, num := range numbers {
		if num%2 == 0 {
			return num

		}
	}
	return -1
}

// Lint Issue 9: Unnecessary else
func CheckValue(x int) string {
	if x > 10 {
		return "large"
	} else { // Unnecessary else
		return "small"
	}
}

// Lint Issue 10: Non-idiomatic error check
func ProcessFile(filename string) error {
	// Non-idiomatic: assigning and checking error on same line
	if _, err := openFile(filename); err != nil {
		return err
	}

	// Should be separate lines for declaration and check
	var data []byte
	var readErr error
	data, readErr = readData(filename)
	if readErr != nil {
		return readErr
	}

	return process(data)
}

// Lint Issue 11: Type assertion without check
func HandleInterface(val interface{}) {
	str := val.(string) // Dangerous: could panic
	fmt.Println(str)
}

// Lint Issue 12: defer in loop
func ProcessFiles(filenames []string) error {
	for _, filename := range filenames {
		file, err := openFile(filename)
		if err != nil {
			return err
		}
		// defer in loop - could cause resource leak
		// Fixed by closing explicitly
		func() {
			defer func() { _ = file.Close() }()
		}()

		// Process file
	}
	return nil
}

// Lint Issue 13: String literal duplication
func CreateMessages(name string) (string, string, string) {
	greeting := "Hello, " + name + "! Welcome to our service."            // Duplicated string
	warning := "Hello, " + name + "! Please check your settings."         // Duplicated string
	farewell := "Goodbye, " + name + "! Thank you for using our service." // Similar pattern

	return greeting, warning, farewell
}

// Lint Issue 14: Function returns are not at top level
func ComplexReturn(x int) (result int, err error) {
	if x < 0 {
		result = 0
		err = fmt.Errorf("negative value")
		return
	}

	// Multiple calculations
	y := x * 2
	z := y + 10

	if z > 100 {
		result = 100
		return
	}

	// More logic
	for i := 0; i < 5; i++ {
		if i == 3 {
			result = z * i
			return
		}
	}

	result = z
	return
}

// Helper functions to avoid compilation errors
type DummyFile struct{}

func (d *DummyFile) Close() error { return nil }

func openFile(name string) (*DummyFile, error) { return &DummyFile{}, nil }
func readData(name string) ([]byte, error)     { return []byte{}, nil }
func process(data []byte) error                { return nil }
