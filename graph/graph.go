package graph

import "fmt"

const (
	// Undiscovered denotes vertex which has not been discovered
	Undiscovered = "UNDISCOVERED"
	// Discovered denotes vertex which has been discovered but not all its edges
	Discovered = "DISCOVERED"
	// Processed denotes vertex which has been processed
	Processed = "PROCESSED"
	// Null denotes empty value
	Null = -1
)

type EdgeNote struct {
	val  int
	next *EdgeNote
}

type Graph struct {
	Edges    map[int]*EdgeNote
	Degree   map[int]int
	M        int // number of edges
	N        int // number of vertices
	Directed bool
}

func NewGraph(n int, directed bool) Graph {
	return Graph{
		Edges:    make(map[int]*EdgeNote),
		Degree:   make(map[int]int),
		M:        0,
		N:        n,
		Directed: directed,
	}
}

func (g *Graph) AddEdge(x int, y int) {
	g.addEdge(x, y)

	if g.Directed == false {
		g.addEdge(y, x)
	}

	g.M++
}

func (g *Graph) addEdge(u int, v int) {
	node := EdgeNote{v, nil}
	// Assign new edge node to beginning of edge list
	if vertex, ok := g.Edges[u]; ok == true {
		node.next = vertex
	}

	g.Edges[u] = &node
	g.Degree[u]++
}

func (g Graph) Display() {
	for k, v := range g.Edges {
		fmt.Printf("%v: ", k)
		for v != nil {
			fmt.Printf(" %v", v.val)
			v = v.next
		}
		fmt.Println("")
	}
}

func (g Graph) CheckEdge(u int, v int) bool {
	if vertex, ok := g.Edges[u]; ok == true {
		for vertex != nil {
			if vertex.val == v {
				return true
			}
			vertex = vertex.next
		}
	}
	return false
}

// FindPath returns path from x -> y if this path exists
func FindPath(x int, y int, parent map[int]int) []int {
	r := findPath(x, y, parent, make([]int, 0))
	if x == y || (len(r) != 0 && r[0] != x) {
		// Check if there is a path from x -> y in parent tree
		return []int{}
	}
	return r
}

func findPath(x int, y int, parent map[int]int, r []int) []int {
	r = append([]int{y}, r...)
	// y == Null returns true when x can not go to y or x == y from the beginning
	if x == y || y == Null {
		return r
	}
	return findPath(x, parent[y], parent, r)
}

