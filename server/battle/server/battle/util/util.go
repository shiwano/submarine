package util

import (
	"math"
)

var floatEpsilon = 0.000001

// EqualFloats determines whether the given floating point numbers are same.
func EqualFloats(a, b float64) bool {
	return math.Abs(a-b) < floatEpsilon
}
