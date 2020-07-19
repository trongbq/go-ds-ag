package sorting

// BubbleSort sorts array by bubbling largest element to top
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ { // Just do inside loop n-1 times so arr[1:] is sorted
		for j := 0; j < n-i-1; j++ { // Start bubble largest element to n-i-1 position
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
