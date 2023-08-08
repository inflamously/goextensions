package collections

import "testing"

func TestMinMaxInt(t *testing.T) {
	var numbers = []int{
		1,
		5,
		50,
		10,
		200,
		13,
	}

	a, b := MinMaxInt(numbers)
	if a != 1 {
		t.Fatalf("A != 1")
	}
	if b != 200 {
		t.Fatalf("B != 200")
	}
}
