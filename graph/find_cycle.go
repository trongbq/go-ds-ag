package graph

// FindCycleInGraph find any cycles in graph, works similar as in DFS
func FindCycleInGraph(g Graph, start int) [][]int {
	parent := make(map[int]int)
	status := make(map[int]string)

	// Initialize data
	for i := 1; i <= g.N; i++ {
		status[i] = Undiscovered
	}

	parent[start] = Null
	status[start] = Discovered

	return findCycle(g, start, &parent, &status)
}

func findCycle(g Graph, u int, parent *map[int]int, status *map[int]string) [][]int {
	cycles := make([][]int, 0)

	v := g.Edges[u] // v is the head of linked list of u's connected vertices
	for v != nil {
		if (*status)[v.val] == Undiscovered {
			(*status)[v.val] = Discovered
			(*parent)[v.val] = u
			cs := findCycle(g, v.val, parent, status)
			if len(cs) != 0 {
				cycles = append(cycles, cs...)
			}
		}
		// If v is discovered and either v is not parent of u (edge is tree edge) or v is parent of
		// u but this is directed graph so there might be two edges in reverse direction
		if (*status)[v.val] == Discovered && ((*parent)[u] != v.val || g.Directed) {
			// Found a cycle
			cycles = append(cycles, FindPath(v.val, u, *parent))
		}
		v = v.next
	}
	(*status)[u] = Processed

	return cycles
}
