package sorting

// MergeSort sorts array using divide and conquer approach
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l int, h int) {
	if l < h {
		m := (h + l) / 2 // Divide into two arrays: [l;m] and (m:h]
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, h)
		merge(arr, l, m, h)
	}
}

func merge(arr []int, l int, m int, h int) {
	n1 := m - l + 1
	n2 := h - m
	// Copy two sub-arrays into two array buffer -> make room for final single sorted array
	// Works better with linked list since we don't have to use buffers
	b1 := make([]int, n1, n1)
	b2 := make([]int, n2, n2)
	copy(b1, arr[l:m+1])
	copy(b2, arr[m+1:h+1])

	i, j, k := 0, 0, l
	for i != n1 || j != n2 {
		if j == n2 || (i != n1 && b1[i] < b2[j]) {
			arr[k] = b1[i]
			k++
			i++
			continue
		}
		arr[k] = b2[j]
		k++
		j++
	}
}
