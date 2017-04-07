package lmath

import (
	"fmt"
	"math"
)

// Vector2 is a 2D vector as in coordinate space (a point)
type Vector2 struct {
	X, Y float64
}

func (m Vector2) Equals(n Vector2) bool {
	return (math.Abs(m.X - n.X) + math.Abs(m.Y - n.Y)) == 0
}
