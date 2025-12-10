package utils

import "strconv"

// ParseNumbers converts string arguments to float64
func ParseNumbers(a, b string) (float64, float64) {
	num1, _ := strconv.ParseFloat(a, 64)
	num2, _ := strconv.ParseFloat(b, 64)
	return num1, num2
}

// IsEven returns true if number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// ReverseString returns reversed string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
