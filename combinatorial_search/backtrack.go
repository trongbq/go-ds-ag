package comsearch

// Subset returns all combination of integers range from 0 to n
func Subset(n int) [][]int {
	c := [2]bool{true, false}
	arr := make([]bool, n+1)

	return subset(0, n, arr, c)
}

func subset(k int, n int, arr []bool, c [2]bool) [][]int {
	if k == n {
		temp := make([]int, 0)
		for i := 0; i <= n; i++ {
			if arr[i] == true {
				temp = append(temp, i)
			}
		}
		return [][]int{temp}
	}

	k++
	var r [][]int
	for i := 0; i < len(c); i++ {
		arr[k] = c[i]
		temp := subset(k, n, arr, c)
		if len(temp) != 0 {
			r = append(r, temp...)
		}
	}
	return r
}
