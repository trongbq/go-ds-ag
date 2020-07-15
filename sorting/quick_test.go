package sorting

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{88, 6, 3, 2, 1, 84, 15, 9852, 1253, 41, 7}
	expected := []int{1, 2, 3, 6, 7, 15, 41, 84, 88, 1253, 9852}

	t.Log("Given an unsorted array")
	{
		t.Logf("\tTest 1:\t When using quick sort to sort the array")
		QuickSort(arr)
		if !Equal(arr, expected) {
			t.Errorf("\t%s\tTest 1:\tShould get the array sorted increasingly", failed)
			t.Logf("\t\t\t Expected: %v", expected)
			t.Logf("\t\t\t Actual: %v", arr)
		} else {
			t.Logf("\t%s\tTest 1: Should get the array sorted increasingly", succeed)
		}
	}
}
