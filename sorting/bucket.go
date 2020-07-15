package sorting

// BucketSort sorts array with uniformly distributed over a range
func BucketSort(arr []float64) {
	n := len(arr)
	b := make(map[int][]float64)
	// Create a bucket
	for i := 0; i < n; i++ {
		idx := int(float64(n) * arr[i])
		if v, ok := b[idx]; ok == true {
			b[idx] = append(v, arr[i])
			continue
		}
		b[idx] = []float64{arr[i]}
	}
	for k, v := range b {
		// Use selection sort to sort values in same key
		SelectionSortFloat(v)
		b[k] = v
	}
	// Withdaw all value from bucket to temporary array
	out := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, b[i]...)
	}
	// Copy back to original array
	for i := 0; i < n; i++ {
		arr[i] = out[i]
	}
}
