package graph

// FordFulkerson return maximum flow value from `s` to `t`
func FordFulkerson(g [][]int, s int, t int) int {
	mf := 0
	rg := make([][]int, len(g))
	for i := 0; i < len(g); i++ {
		rg[i] = make([]int, len(g[i]))
		for j := 0; j < len(g[i]); j++ {
			rg[i][j] = g[i][j]
		}
	}

	var parent []int
	for true {
		parent = BfsMatrix(rg, s)
		if parent[t] == Null { // Can not reach t from s
			break
		}

		f := MaxInt
		// Find min flow from s to t
		for v := t; v != s; v = parent[v] {
			u := parent[v]
			if rg[u][v] < f {
				f = rg[u][v]
			}
		}

		// Update residual graph rg with min flow
		for v := t; v != s; v = parent[v] {
			u := parent[v]
			rg[u][v] -= f
			rg[v][u] += f // To allow redo
		}

		mf += f
	}

	return mf
}
