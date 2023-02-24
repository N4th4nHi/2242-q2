//go interface
package main

import (
	"fmt"
)

type mathOperations interface {
	Add(a, b float64) float64
	subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) float64
}

type Math struct{}

func (m Math) Add(a, b float64) float64 {
	return a + b
}

func (m Math) Subtract(a, b float64) float64 {
	return a - b
}

func (m Math) Multiply(a, b float64) float64 {
	return a * b
}

func (m Math) Divide(a, b float64) float64 {
	return a / b
}

func main() {
	math := Math{}
	a := 5.0
	b := 2.0

	fmt.Println(a, "+", b, "=", math.Add(a, b))
	fmt.Println(a, "-", b, "=", math.Subtract(a, b))
	fmt.Println(a, "*", b, "=", math.Multiply(a, b))
	fmt.Println(a, "/", b, "=", math.Divide(a, b))
}
