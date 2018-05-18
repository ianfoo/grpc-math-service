package mather

import "testing"

func TestAdd(t *testing.T) {
	tt := []struct {
		desc     string
		a, b     int64
		expected int64
	}{
		{desc: "zeros", a: 0, b: 0, expected: 0},
		{desc: "ones", a: 1, b: 1, expected: 2},
		{desc: "regular", a: 20, b: 22, expected: 42},
		{desc: "large", a: 1.2e10, b: 3.9e10, expected: 5.1e10},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			if actual := Add(tc.a, tc.b); actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tt := []struct {
		desc     string
		a, b     int64
		expected int64
	}{
		{desc: "zeros", a: 0, b: 0, expected: 0},
		{desc: "ones", a: 1, b: 1, expected: 0},
		{desc: "regular", a: 173, b: 131, expected: 42},
		{desc: "large", a: 1.8e10, b: 3.3e10, expected: -1.5e10},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			if actual := Subtract(tc.a, tc.b); actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tt := []struct {
		desc     string
		a, b     int64
		expected int64
	}{
		{desc: "zeros", a: 0, b: 0, expected: 0},
		{desc: "ones", a: 1, b: 1, expected: 1},
		{desc: "regular", a: 14, b: 3, expected: 42},
		{desc: "large", a: 1.2e7, b: 3.9e7, expected: 4.68e14},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			if actual := Multiply(tc.a, tc.b); actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tt := []struct {
		desc              string
		a, b              int64
		expectedQuotient  int64
		expectedRemainder int64
		expectedError     error
	}{
		{desc: "zeros", a: 0, b: 0, expectedQuotient: 0, expectedRemainder: 0, expectedError: ErrDivideByZero},
		{desc: "ones", a: 1, b: 1, expectedQuotient: 1, expectedRemainder: 0},
		{desc: "regular", a: 173, b: 131, expectedQuotient: 1, expectedRemainder: 42},
		{desc: "large", a: 1.2e10, b: 3.9e10, expectedQuotient: 0, expectedRemainder: 1.2e10},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			if q, r, err := Divide(tc.a, tc.b); err != tc.expectedError {
				t.Errorf("expected error %v but got %v", tc.expectedError, err)
			} else if q != tc.expectedQuotient || r != tc.expectedRemainder {
				t.Errorf("expected %d quotient and %d remainder, "+
					"but got %d quotient and %d remainder",
					tc.expectedQuotient, tc.expectedRemainder, q, r)
			}
		})
	}
}
