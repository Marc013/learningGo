//go:build solution

package task2

import (
	"os"
	"testing"
)

var c *Calculator

func TestMain(m *testing.M) {
	c = &Calculator{}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestAdd(t *testing.T) {
	t.Parallel()

	c.Reset(1)

	result := c.Add(2)

	expected := 3.0
	if result != expected {
		t.Errorf("Addition was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	c.Reset(1)

	result := c.Subtract(3)

	expected := -2.0
	if result != expected {
		t.Errorf("Subtraction was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	c.Reset(1)

	result := c.Multiply(5)

	expected := 5.0
	if result != expected {
		t.Errorf("Multiplication was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestDivideValid(t *testing.T) {
	t.Parallel()

	c.Reset(1)

	result, err := c.Divide(7)

	if err != nil {
		t.Errorf("Error while dividing: %v", err)
	}

	expected := 1.0 / 7.0
	if result != expected {
		t.Errorf("Division was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	t.Parallel()

	c.Reset(1)

	_, err := c.Divide(0)

	if err == nil {
		t.Error("Expected error for division by zero but got none")
	}
}
