package sorting

import "testing"

func TestCountingSort(t *testing.T) {
	// arr := []int{3, 1, 5, 8, 1, 2, 6, 9, 7, 3}
	// expected := []int{1, 1, 2, 3, 3, 5, 6, 7, 8, 9}
	arr := []rune("todayt")
	expected := []rune{'a', 'd', 'o', 't', 't', 'y'}
	t.Log("Given an unsorted array")
	{
		t.Logf("\tTest 1:\t When using counting sort to sort the array")
		CountingSort(arr)
		if !EqualRuneArray(arr, expected) {
			t.Errorf("\t%s\tTest 1:\tShould get the array sorted increasingly", failed)
			t.Logf("\t\t\t Want: %v", string(expected))
			t.Logf("\t\t\t Got: %v", string(arr))
		} else {
			t.Logf("\t%s\tTest 1: Should get the array sorted increasingly", succeed)
		}
	}
}
