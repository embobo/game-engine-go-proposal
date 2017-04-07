package gmath

import (
	"fmt"
	"math"
)

// Vector2 is a 2D vector as in coordinate space (a point)
type Vector2 struct {
	X, Y float64
}

// Componentwise equals
func (m Vector2) Equals(n Vector2) bool {
	return (math.Abs(m.X - n.X) + math.Abs(m.Y - n.Y)) == 0
}

// Copies vector
func (m Vector2) Copy() Vector2 {
	return Vector2{m.X, m.Y}
}

// Componentwise addition
func (m Vector2) Add(n Vector2) Vector2 {
	return Vector2{m.X + n.X, m.Y + n.Y}
}

// Componentwise subtraction
func (m Vector2) Subtract(n Vector2) Vector2 {
	return Vector2{m.X - n.X, m.Y - n.Y}
}

// Componentwise multiplication
func (m Vector2) Multiply(n Vector2) Vector2 {
	return Vector2{m.X * n.X, m.Y * n.Y}
}

// Componentwise division
func (m Vector2) Divide(n Vector2) Vector2 {
	return Vector2{m.X / n.X, m.Y / n.Y}
}

// Componentwise scalar addition
func (m Vector2) AddScalar(k float64) Vector2 {
	return Vector2{m.X + k, m.Y + k}
}

// Componentwise scalar subtraction
func (m Vector2) SubtractScalar(k float64) Vector2 {
	return Vector2{m.X - k, m.Y - k}
}

// Componentwise scalar multiplication
func (m Vector2) MultiplyScalar(k float64) Vector2 {
	return Vector2{m.X * k, m.Y * k}
}

// Componentwise scalar division
func (m Vector2) DivideScalar(k float64) Vector2 {
	return Vector2{m.X / k, m.Y / k}
}

// Gets Magnitude
func (m Vector2) Magnitude() float64 {
	return Sqrt(Pow(m.X, 2) + Pow(m.Y, 2))
}

// Gets Magnitude of m < n
func (m Vector2) LessThan(n Vector2) bool {
	return m.Magnitude() < n.Magnitude()
}

// Gets Magnitude of m > n
func (m Vector2) GreaterThan(n Vector2) bool {
	return m.Magnitude() > n.Magnitude()
}
