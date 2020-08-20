package graph

const (
	// MaxUint is maximum unsigned integer
	MaxUint = ^uint(0)
	// MaxInt is maximum signed integer
	MaxInt = int(MaxUint >> 1)
)

// Prim follow Prim algorithm to find minimum spanning tree and return parent array (for path construction)
func Prim(g Graph, start int) map[int]int {
	parent := make(map[int]int)

	// Holds information for all vertices on whether they are in spanning tree
	// Avoid using index 0 to simplify implementation
	in := make([]bool, g.N+1)
	distance := make([]int, g.N+1)
	for i := 1; i <= g.N; i++ {
		in[i] = false
		distance[i] = MaxInt
		parent[i] = Null
	}

	u := start
	distance[u] = 0

	for u != Null {
		in[u] = true

		v := g.Edges[u]
		for v != nil {
			if !in[v.val] && v.weight < distance[v.val] {
				distance[v.val] = v.weight
				parent[v.val] = u
			}
			v = v.next
		}

		// Find next vertex to join minimum spanning tree
		u = Null
		dist := MaxInt
		for i := 1; i < len(distance); i++ {
			if !in[i] && dist > distance[i] {
				u = i
				dist = distance[i]
			}
		}
	}

	return parent
}
