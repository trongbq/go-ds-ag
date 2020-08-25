package comsearch

// Subset returns all combination of integers range from 0 to n
func Subset(n int) [][]int {
	c := [2]bool{true, false}
	arr := make([]bool, n+1)

	return subset(-1, n, arr, c)
}

// Permuation return all permutation from 0 to n
func Permuation(n int) [][]int {
	arr := make([]int, n+1)
	return perm(-1, n, arr)
}

// CountAllPathInGraph returns number of path that can go from s to t
func CountAllPathInGraph(g [][]int, s, t int) int {
	arr := make([]int, len(g))
	arr[0] = s
	return countPath(g, 0, t, arr)
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

func perm(k, n int, arr []int) [][]int {
	if k == n {
		return [][]int{arr}
	}

	// Store memory about all values which is selected from previous operation
	incl := make([]bool, n+1)
	for i := 0; i <= k; i++ {
		incl[arr[i]] = true
	}

	// Construct all candiates which are unselected numbers
	var c []int
	for i := 0; i <= n; i++ {
		if !incl[i] {
			c = append(c, i)
		}
	}

	k++
	var r [][]int
	for _, v := range c {
		arr[k] = v
		temp := perm(k, n, arr)
		if len(temp) != 0 {
			r = append(r, temp...)
		}
	}
	return r
}

func countPath(g [][]int, k, t int, arr []int) int {
	if arr[k] == t {
		return 1
	}

	// Store memory about all values which is selected from previous operation
	// This will avoid choosing a circle
	incl := make([]bool, len(g))
	for i := 0; i <= k; i++ {
		incl[arr[i]] = true
	}

	// Construct all candiates which are connected to vertex k
	var c []int
	for i := 0; i < len(g); i++ {
		if g[arr[k]][i] > 0 && !incl[i] {
			c = append(c, i)
		}
	}

	k++
	var count int
	for _, v := range c {
		arr[k] = v
		temp := countPath(g, k, t, arr)
		count += temp
	}
	return count
}
