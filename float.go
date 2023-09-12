package math32

import (
	"math"
)

// Mathematical constants.
const (
	Pi float32 = math.Pi
)

// PosInf is a floating-point representation of positive infinity.
var PosInf = math.Float32frombits(0x7F800000)

// NegInf is a floating-point representation of negative infinity.
var NegInf = math.Float32frombits(0xFF800000)

// Abs returns the absolute value of x.
func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

// Clamp returns the closest value to x within the range [min, max].
func Clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Cos returns the cosine of the radian argument x.
func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

// Max returns the larger of x or y.
func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
func Min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

// Sign returns +1 or -1 depending on the sign of x.
func Sign(x float32) float32 {
	if x < 0 {
		return -1
	}
	return 1
}

// Sin returns the sine of the radian argument x.
func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

// Sqrt returns the square root of x.
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
