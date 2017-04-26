package gmath

import "testing"

func TestEqual(test *testing.T) {
	v := Vector2{1,2}
	if !v.Equals(Vector2{1,2}) {
		test.Fail()
	}
	if v.Equals(Vector2{1,0}) {
		test.Fail()
	}
}

func TestAdd(test *testing.T) {
	v := Vector2{1,2}
	r := v.Add(Vector2{2,3})
	if !(r.Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestSelfAdd(test *testing.T) {
	v := Vector2{1,2}
	r := Vector2{2,3}
	v.SelfAdd(r)
	if !(v.Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestSubtract(test *testing.T) {
	v := Vector2{1,2}
	r := v.Subtract(Vector2{2,2})
	if !(r.Equals(Vector2{-1,0})) {
		test.Fail()
	}
}

func TestSelfSubtract(test *testing.T) {
	v := Vector2{3,5}
	v.SelfSubtract(Vector2{4,2})
	if !(v.Equals(Vector2{-1,3})) {
		test.Fail()
	}
}

func TestMultiply(test *testing.T) {
	v := Vector2{1,2}
	r := v.Multiply(Vector2{4,2})
	if !(r.Equals(Vector2{4,4})) {
		test.Fail()
	}
}

func TestSelfMultiply(test *testing.T) {
	v := Vector2{1,2}
	v.SelfMultiply(Vector2{4,-2})
	if !(v.Equals(Vector2{4,-4})) {
		test.Fail()
	}
}

func TestDivide(test *testing.T) {
	v := Vector2{3,8}
	r := v.Divide(Vector2{2,4})
	if !(r.Equals(Vector2{1.5,2})) {
		test.Fail()
	}
}

func TestSelfDivide(test *testing.T) {
	v := Vector2{4,4}
	v.SelfDivide(Vector2{2,2})
	if !(v.Equals(Vector2{2,2})) {
		test.Fail()
	}
}

func TestAddScalar(test *testing.T) {
	v := Vector2{1,3}
	r := v.AddScalar(2)
	if !(r.Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestSelfAddScalar(test *testing.T) {
	v := Vector2{1,3}
	v.SelfAddScalar(2)
	if !(v.Equals(Vector2{3,5})) {
		test.Fail()
	}
}

func TestSubtractScalar(test *testing.T) {
	v := Vector2{1,2}
	r := v.SubtractScalar(2)
	if !(r.Equals(Vector2{-1,0})) {
		test.Fail()
	}
}

func TestSelfSubtractScalar(test *testing.T) {
	v := Vector2{1,3}
	v.SelfSubtractScalar(2)
	if !(v.Equals(Vector2{-1,1})) {
		test.Fail()
	}
}

func TestMultiplyScalar(test *testing.T) {
	v := Vector2{1,2}
	r := v.MultiplyScalar(3)
	if !(r.Equals(Vector2{3,6})) {
		test.Fail()
	}
}

func TestDivideScalar(test *testing.T) {
	v := Vector2{3,8}
	r := v.DivideScalar(2)
	if !(r.Equals(Vector2{1.5,4})) {
		test.Fail()
	}
}

func TestMagnitude(test *testing.T) {
	v := Vector2{3,4}
	if !(v.Magnitude() == 5) {
		test.Fail()
	}
}

func TestLessThan(test *testing.T) {
	v := Vector2{1,2}
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

func TestGreaterThan(test *testing.T) {
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
