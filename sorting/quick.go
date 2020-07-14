package sorting

// QuickSort ...
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l int, h int) {
	if l < h {
		p := partition(arr, l, h)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, h)
	}
}

// partition finds the pivot index `p` that
// for all elements in arr[:p] < arr[p] and elements in arr[p+1:] > p
func partition(arr []int, l int, h int) int {
	p := l
	for i := l; i < h; i++ {
		// Take the element at position `h` be the pivot element
		// If any element less than pivot element, swap it with current pivot index `p`
		if arr[i] < arr[h] {
			arr[i], arr[p] = arr[p], arr[i]
			p++
		}
	}
	// Swap pivot to its correct index
	arr[p], arr[h] = arr[h], arr[p]
	return p
}
