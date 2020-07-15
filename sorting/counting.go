package sorting

// MaxCharRange rune scope size is 26 for 26 characters
const MaxCharRange = 26

// CountingSort sort array with key serves as index in array
func CountingSort(arr []rune) {
	count := make([]int, MaxCharRange, MaxCharRange)
	for _, v := range arr {
		count[v-'a']++
	}
	// count[i] contains number of elements exists in the original array which are less than or equal to `i`
	for i := 1; i < MaxCharRange; i++ {
		count[i] = count[i] + count[i-1]
	}
	out := make([]rune, len(arr), len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		out[count[arr[i]-'a']-1] = arr[i]
		count[arr[i]-'a']--
	}
	// Copy back to origin array
	for i := 0; i < len(arr); i++ {
		arr[i] = out[i]
	}
}
