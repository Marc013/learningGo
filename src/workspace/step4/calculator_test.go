package step4

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := Calculator{}

	expected := 5.0

	result := c.Add(2, 3)

	if result != expected {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	c := Calculator{}

	expected := 2.0

	result := c.Subtract(5, 3)

	if result != expected {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	c := Calculator{}

	expected := 6.0

	result := c.Multiply(2, 3)

	if result != expected {
		t.Fail()
	}
}

func TestDivideValid(t *testing.T) {
	c := Calculator{}

	expected := 2.0

	result, err := c.Divide(6, 3)

	if err != nil {
		t.Fail()
	}

	if result != expected {
		t.Fail()
	}
}

func TestDivideByZero(t *testing.T) {
	c := Calculator{}

	_, err := c.Divide(6, 0)

	if err == nil {
		t.Fail()
	}
}
