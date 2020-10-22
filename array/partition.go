package array

func DutchFlagPartition(pivotIndex int, arr []int) {
	pivot := arr[pivotIndex]
	smaller := 0 // hold the next position of smaller elements
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			arr[smaller], arr[i] = arr[i], arr[smaller]
			smaller++
		}
	}
	larger := len(arr) - 1
	for i := len(arr) - 1; i >= 0 && arr[i] >= pivot; i-- {
		if arr[i] > pivot {
			arr[larger], arr[i] = arr[i], arr[larger]
			larger--
		}
	}
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
