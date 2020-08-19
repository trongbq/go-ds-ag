package graph

const (
	TreeEdge = "TREE_EDGE"
	BackEdge = "TREE_EDGE"
	ForwardEdge = "TREE_EDGE"
	CrossEdge = "TREE_EDGE"
)
type Record struct {
	In  int
	Out int
}

// Dfs runs depth-first search through undirected graph and produces map parent of vertices
func Dfs(g Graph, start int) (map[int]int, map[int]Record) {
	parent := make(map[int]int)
	records := make(map[int]Record)
	status := make(map[int]string)

	// Initialize data
	for i := 1; i <= g.N; i++ {
		status[i] = Undiscovered
	}

	parent[start] = Null
	status[start] = Discovered
	t := 0

	// Using recursive as a stack of vertices
	dfs(g, start, &t, &parent, &records, &status)

	return parent, records
}

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

func edgeClassification(u int, v int, status map[int]string, parent map[int]int, record map[int]Record) string {
	if parent[v] == u {
		return TreeEdge
	}
	if status[v] == Discovered {
		return BackEdge
	}
	if status[v] == Processed && record[v].In > record[u].In {
		return ForwardEdge
	}
	if status[v] == Processed && record[v].In < record[u].In {
		return CrossEdge
	}

	return ""
}

func dfs(g Graph, u int, t *int, parent *map[int]int, records *map[int]Record, status *map[int]string) {
	(*t)++
	r := Record{In: *t}

	v := g.Edges[u] // v is the head of linked list of u's connected vertices
	for v != nil {
		if (*status)[v.val] == Undiscovered {
			(*status)[v.val] = Discovered
			(*parent)[v.val] = u
			dfs(g, v.val, t, parent, records, status)
		}
		v = v.next
	}

	(*t)++
	r.Out = *t
	(*records)[u] = r

	(*status)[u] = Processed
}

