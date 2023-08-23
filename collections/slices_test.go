package collections

import (
	"reflect"
	"testing"
)

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

func TestReverse(t *testing.T) {
	input := []int{1, 3, 5, 6}
	Reverse(input)
	output := []int{6, 5, 3, 1}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("Input does not match Output!")
	}
}

func TestReverse_2(t *testing.T) {
	input := []int{1, 3, 5, 4, 6}
	Reverse(input)
	output := []int{6, 4, 5, 3, 1}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("Input does not match Output!")
	}
}

func TestReverse_3(t *testing.T) {
	input := []int{1, 3}
	Reverse(input)
	output := []int{3, 1}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("Input does not match Output!")
	}
}
