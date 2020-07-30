package graph

import "fmt"

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
