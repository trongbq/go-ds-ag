package graph

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
