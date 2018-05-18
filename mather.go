package mather

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

// Add returns the sum of its two arguments.
func Add(a, b int64) int64 {
	return a + b
}

// Subtract returns the difference of its two arguments, where b is subtracted
// from a.
func Subtract(a, b int64) int64 {
	return a - b
}

// Multiply returns the product of its two arguments.
func Multiply(a, b int64) int64 {
	return a * b
}

// Divide returns the quotient and remainder of its two arguments, where a is
// the divident, and b is the divisor.
func Divide(a, b int64) (int64, int64, error) {
	if b == 0 {
		return 0, 0, ErrDivideByZero
	}
	return a / b, a % b, nil
}
