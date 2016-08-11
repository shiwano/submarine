package util

import (
	"math"
)

// EqualFloats determines whether the given floating point numbers are same.
func EqualFloats(a, b float64) bool {
	return math.Abs(a-b) < math.SmallestNonzeroFloat64
}
