package state

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	value *TestStruct
	x     string
}

func (s *TestStruct) Clone() TestStruct {
	newValue := *s.value
	return TestStruct{
		x:     s.x,
		value: &newValue,
	}
}

func TestPointerClone(t *testing.T) {
	subValue := &TestStruct{
		value: nil,
		x:     "b",
	}
	var testA *TestStruct = &TestStruct{
		value: subValue,
		x:     "a",
	}

	//storedCopy := *subValue
	var testB *TestStruct = &TestStruct{
		value: ShallowCopy[TestStruct](subValue),
		x:     "a",
	}

	var testC = testA.Clone()

	testB.value.x = "Cushy!"
	testC.x = "Yujup!"

	//fmt.Printf("%p | %p %p", testA, testB, testB.value)

	isEqualTestA := reflect.DeepEqual(testA, testB)
	if isEqualTestA {
		t.Fatalf("testA AND testB should NOT be equal")
	}

	isEqualTestB := reflect.DeepEqual(testA, testC)
	if isEqualTestB {
		t.Fatalf("testA AND testB should NOT be equal")
	}
}
