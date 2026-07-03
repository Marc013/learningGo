package step5

import (
	"errors"
	"time"
)

type Calculator struct {
	value float64
}

func (c *Calculator) Add(a float64) float64 {
	time.Sleep(100 * time.Millisecond)
	c.value += a
	return c.value
}

func (c *Calculator) Subtract(a float64) float64 {
	time.Sleep(100 * time.Millisecond)
	c.value -= a
	return c.value
}

func (c *Calculator) Multiply(a float64) float64 {
	time.Sleep(100 * time.Millisecond)
	c.value *= a
	return c.value
}

func (c *Calculator) Divide(a float64) (float64, error) {
	time.Sleep(100 * time.Millisecond)
	if a == 0 {
		return 0, errors.New("division by zero")
	}
	c.value /= a
	return c.value, nil
}

func (c *Calculator) Reset(a float64) {
	c.value = a
}
