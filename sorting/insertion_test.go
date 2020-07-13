package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{88, 6, 3, 2, 1, 84, 15, 9852, 1253, 41, 7}
	expected := []int{1, 2, 3, 6, 7, 15, 41, 84, 88, 1253, 9852}
	InsertionSort(arr)
	if !Equal(arr, expected) {
		t.Errorf("InsertionSort failed\nExpected:%v\nActual:%v", expected, arr)
	}
}
