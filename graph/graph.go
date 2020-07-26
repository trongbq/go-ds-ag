package graph

import "fmt"

type EdgeNote struct {
	y    int
	next *EdgeNote
}

type Graph struct {
	edges    map[int]*EdgeNote
	degree   map[int]int
	m        int // number of edges
	n        int // number of vertices
	directed bool
}

func NewGraph(n int, directed bool) *Graph {
	return &Graph{
		edges:    make(map[int]*EdgeNote),
		degree:   make(map[int]int),
		m:        0,
		n:        n,
		directed: directed,
	}
}

func (g *Graph) AddEdge(x int, y int) {
	g.addEdge(x, y)

	if g.directed == false {
		g.addEdge(y, x)
	}

	g.m++
}

func (g *Graph) addEdge(x int, y int) {
	node := EdgeNote{y, nil}
	// Assign new edge node to beginning of edge list
	if v, ok := g.edges[x]; ok == true {
		node.next = v
	}

	g.edges[x] = &node
	g.degree[x]++
}

func (g Graph) Display() {
	for k, v := range g.edges {
		fmt.Printf("%v: ", k)
		for v != nil {
			fmt.Printf(" %v", v.y)
			v = v.next
		}
		fmt.Println("")
	}
}

func (g Graph) CheckEdge(x int, y int) bool {
	if v, ok := g.edges[x]; ok == true {
		for v != nil {
			if v.y == y {
				return true
			}
			v = v.next
		}
	}
	return false
}
