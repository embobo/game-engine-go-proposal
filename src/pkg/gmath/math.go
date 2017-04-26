package gmath

//import "math"

func Equal(x, y float64) bool {
	return x == y
}

// Lerp (0, 2, 0) = 0
// Lerp (0, 2, 0.5) = 1
// Lerp (0, 2, 1) = 2

func Lerp(x, y, t float64) float64 {
	return (1-t)*x + t*y
}
