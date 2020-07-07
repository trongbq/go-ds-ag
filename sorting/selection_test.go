package sorting

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{88, 6, 3, 2, 1, 84, 15, 9852, 1253, 41, 7}
	expected := []int{1, 2, 3, 6, 7, 15, 41, 84, 88, 1253, 9852}
	SelectionSort(arr)
	if !Equal(arr, expected) {
		t.Errorf("SelectionSort failed\nExpected:%v\nActual:%v", expected, arr)
	}
}

func TestSelectionSortDupplicate(t *testing.T) {
	arr := []int{88, 6, 3, 2, 1, 84, 15, 9852, 1253, 41, 7, 7, 7}
	expected := []int{1, 2, 3, 6, 7, 7, 7, 15, 41, 84, 88, 1253, 9852}
	SelectionSort(arr)
	if !Equal(arr, expected) {
		t.Errorf("SelectionSort failed\nExpected:%v\nActual:%v", expected, arr)
	}
}
