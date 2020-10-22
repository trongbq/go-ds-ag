package array

import "testing"

func TestDutchFlagPartition(t *testing.T) {
	input := []int{0, 1, 2, 0, 2, 1, 1}
	pivotIndex := 1
	pivot := input[pivotIndex]
	DutchFlagPartition(pivotIndex, input)
	i := 0
	for i = 0; i < len(input) && input[i] < pivot; i++ {
	}
	for ; i < len(input) && input[i] == pivot; i++ {
	}
	for ; i < len(input) && input[i] > pivot; i++ {
	}
	if i != len(input) {
		t.Fatalf("DutchFlagPartition test failed, got: %v", input)
	}
}
