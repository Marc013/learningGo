//go:build solution

package task3

import (
	"os"
	"testing"
)

var c Calculator

func TestMain(m *testing.M) {
	c = Calculator{}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestCalculator(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		expected := 5.0

		result := c.Add(2, 3)

		if result != expected {
			t.Errorf("Addition was incorrect, got: %f, want: %f.", result, expected)
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		expected := 2.0

		result := c.Subtract(5, 3)

		if result != expected {
			t.Errorf("Subtraction was incorrect, got: %f, want: %f.", result, expected)
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		expected := 6.0

		result := c.Multiply(2, 3)

		if result != expected {
			t.Errorf("Multiplication was incorrect, got: %f, want: %f.", result, expected)
		}
	})

	t.Run("Divide", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			expected := 2.0

			result, err := c.Divide(6, 3)

			if err != nil {
				t.Errorf("Error while dividing: %v", err)
			}

			if result != expected {
				t.Errorf("Division was incorrect, got: %f, want: %f.", result, expected)
			}
		})

		t.Run("ByZero", func(t *testing.T) {
			_, err := c.Divide(6, 0)

			if err == nil {
				t.Error("Expected error for division by zero but got none")
			}
		})
	})
}
