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
