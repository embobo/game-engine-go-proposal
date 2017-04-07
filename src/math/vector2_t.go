package lmath

import "testing"

func TestVector2Equal(test *testing.T) {
	v := Vector2{1,2}
	if !v.Equals(Vector2{1,2}) {
		test.Fail()
	}
	if v.Equals(Vector2{1,0}) {
		test.Fail()
	}
}

func TestVector2Add(test *testing.T) {
	v := Vector2{1,2}
	if !(v.Add(Vector2{2,3}).Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestVector2Subtract(test *testing.T) {
	v := Vector2{1,2}
	if !(v.Subtract(Vector2{2,2}).Equals(Vector2{-1,0})) {
		test.Fail()
	}
}

func TestVector2Multiply(test *testing.T) {
	v := Vector2{1,2}
	if !(v.Multiply(Vector2{4,2}).Equals(Vector2{4,4})) {
		test.Fail()
	}
}

func TestVector2Divide(test *testing.T) {
	v := Vector2{3,8}
	if !(v.Divide(Vector2{2,4}).Equals(Vector2{1.5,2})) {
		test.Fail()
	}
}

func TestVector2AddScalar(test *testing.T) {
	v := Vector2{1,3}
	if !(v.AddScalar(2).Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestVector2SubtractScalar(test *testing.T) {
	v := Vector2{1,2}
	if !(v.SubtractScalar(2).Equals(Vector2{-1,0})) {
		test.Fail()
	}
}

func TestVector2MultiplyScalar(test *testing.T) {
	v := Vector2{1,2}
	if !(v.MultiplyScalar(3).Equals(Vector2{3,6})) {
		test.Fail()
	}
}

func TestVector2DivideScalar(test *testing.T) {
	v := Vector2{3,8}
	if !(v.DivideScalar(2).Equals(Vector2{1.5,4})) {
		test.Fail()
	}
}

func TestVector2Magnitude(test *testing.T) {
	v := Vector{3,4}
	if !(v.Magnitude() == 5) {
		test.Fail()
	}
}

func TestVector2LessThan(test *testing.T) {
	v := Vector{1,2}
	if v.LessThan(Vector2{0,2}) {
		test.Fail()
	}
	if v.LessThan(v) {
		test.Fail()
	}
	if !v.LessThan(Vector2{2,3}) {
		test.Fail()
	}
}

func TestVector2GreaterThan(test *testing.T) {
	v := Vector2{1,2}
	if v.GreaterThan(Vector2{1,3}) {
		test.Fail()
	}
	if v.GreaterThan(v) {
		test.Fail()
	}
	if !v.GreaterThan(Vector2{0,2}) {
		test.Fail()
	}
}
