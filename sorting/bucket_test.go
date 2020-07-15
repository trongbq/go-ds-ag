package sorting

import "testing"

func TestBucketSort(t *testing.T) {
	arr := []float64{0.23, 0.55, 0.12, 0.77, 0.43, 0.11, 0.17, 0.52}
	expected := []float64{0.11, 0.12, 0.17, 0.23, 0.43, 0.52, 0.55, 0.77}

	t.Log("Given an unsorted array")
	{
		t.Logf("\tTest 1:\t When using bucket sort to sort the array")
		BucketSort(arr)
		if !EqualFloatArray(arr, expected) {
			t.Errorf("\t%s\tTest 1:\tShould get the array sorted increasingly", failed)
			t.Logf("\t\t\t Want: %v", expected)
			t.Logf("\t\t\t Got: %v", arr)
		} else {
			t.Logf("\t%s\tTest 1: Should get the array sorted increasingly", succeed)
		}
	}
}
