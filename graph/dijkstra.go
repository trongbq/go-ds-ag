package graph

// Dijkstra algorithm return shortest path from `start` vertex to all other vertices
func Dijkstra(g Graph, start int) ([]int, map[int]int) {
	parent := make(map[int]int)
	checked := make([]bool, g.N+1)
	distance := make([]int, g.N+1)
	for i := 1; i <= g.N; i++ {
		checked[i] = false
		distance[i] = MaxInt
		parent[i] = Null
	}

	u := start
	distance[u] = 0

	for u != Null {
		checked[u] = true

		v := g.Edges[u]
		for v != nil {
			// If distance of `v` greater than distance from `start` to `v` though `u`, then update distance of `v`
			if !checked[v.val] && distance[v.val] > (distance[u]+v.weight) {
				distance[v.val] = distance[u] + v.weight
				parent[v.val] = u
			}
			v = v.next
		}

		// Find next shortest vertex to go from `start`
		u = Null
		dist := MaxInt
		for i := 1; i < len(distance); i++ {
			if !checked[i] && dist > distance[i] {
				u = i
				dist = distance[i]
			}
		}
	}

	return distance[1:], parent
}
