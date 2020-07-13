package sorting

// InsertionSort sorts slice increasingly
func InsertionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// arr[:i] is sorted, find the right place for arr[i] so arr[:i+1] is sorted

	// for i := 1; i < len(arr); i++ {
	// 	j := i
	// 	for j > 0 && arr[j-1] > arr[j] {
	// 		arr[j], arr[j-1] = arr[j-1], arr[j]
	// 		j--
	// 	}
	// }
	for i := 1; i < len(arr); i++ {
		j := i
		elem := arr[j]
		for j > 0 && arr[j-1] > elem {
			arr[j] = arr[j-1] // Shift larger element to the right, hold place for `j`
			j--
		}
		arr[j] = elem
	}
}
