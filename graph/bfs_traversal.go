package graph

// Bfs runs bread-first search through undirected graph and produces map parent of vertices
func Bfs(g Graph, start int) map[int]int {
	parent := make(map[int]int)
	status := make(map[int]string)
	var queue []int

	// Initialize data
	for i := 1; i <= g.N; i++ {
		status[i] = Undiscovered
	}

	queue = enqueue(queue, start)
	status[start] = Discovered
	parent[start] = Null

	var u int
	for !isEmpty(queue) {
		u, queue = dequeue(queue)
		v := g.Edges[u] // e is the head of linked list of v's connected vertices
		for v != nil {
			if status[v.val] == Undiscovered {
				queue = enqueue(queue, v.val)
				status[v.val] = Discovered
				parent[v.val] = u
			}
			v = v.next
		}
	}

	return parent
}

func enqueue(q []int, v int) []int {
	return append(q, v)
}

func dequeue(q []int) (int, []int) {
	if len(q) == 0 {
		return Null, q
	}
	temp := q[0]
	q[0] = 0 // Write zero value to avoid memory leaks
	q = q[1:]
	return temp, q
}

func isEmpty(q []int) bool {
	return len(q) == 0
}

