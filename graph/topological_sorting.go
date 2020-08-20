package graph

import (
	"errors"
)

// TopologicalSorting returns order of vertices in DAG
func TopologicalSorting(g Graph) ([]int, error) {
	arr := []int{}
	parent := make(map[int]int)
	status := make(map[int]string)
	records := make(map[int]Record)

	// Initialize data
	for i := 1; i <= g.N; i++ {
		status[i] = Undiscovered
		parent[i] = Null
	}
	t := 0

	for i := 1; i <= g.N; i++ {
		if status[i] == Undiscovered {
			v, err := topologicalSorting(g, i, &t, &parent, &status, &records)
			if err != nil {
				return []int{}, err
			}
			if len(v) != 0 {
				arr = append(arr, v...)
			}
		}
	}

	// Topological sorting result is a reverse order of vertices when they are marked `processed`
	// `arr` is a stack, reverse it
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr, nil
}

func topologicalSorting(g Graph, u int, t *int, parent *map[int]int, status *map[int]string, records *map[int]Record) ([]int, error) {
	arr := []int{}

	(*status)[u] = Discovered

	v := g.Edges[u] // v is the head of linked list of u's connected vertices
	for v != nil {
		if (*status)[v.val] == Undiscovered {
			(*status)[v.val] = Discovered
			(*parent)[v.val] = u
			a, err := topologicalSorting(g, v.val, t, parent, status, records)
			if err != nil {
				return []int{}, err
			}
			arr = append(arr, a...)
		}
		if (*status)[v.val] == Discovered {
			if t := classifyEdge(u, v.val, *status, *parent, *records); t == BackEdge {
				return []int{}, errors.New("not a DAG, cycle found")
			}
		}
		v = v.next
	}

	(*status)[u] = Processed
	arr = append(arr, u)

	return arr, nil
}
