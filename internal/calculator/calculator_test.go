package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("Subtract(5, 3) = %v, want %v", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)
	expected := 12.0
	if result != expected {
		t.Errorf("Multiply(4, 3) = %v, want %v", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		result, err := Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) returned error: %v", err)
		}
		expected := 5.0
		if result != expected {
			t.Errorf("Divide(10, 2) = %v, want %v", result, expected)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) should return error")
		}
	})
}

func TestPower(t *testing.T) {
	result := Power(2, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Power(2, 3) = %v, want %v", result, expected)
	}
}
