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

// FindPath returns shortest path from x -> y if this path exists
func FindPath(x int, y int, parent map[int]int) []int {
	r := findPath(x, y, parent, make([]int, 0))
	if x == y || (len(r) != 0 && r[0] != x) {
		// Check if there is a path from x -> y in parent tree
		return []int{}
	}
	return r
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

func findPath(x int, y int, parent map[int]int, r []int) []int {
	r = append([]int{y}, r...)
	// y == Null returns true when x can not go to y or x == y from the beginning
	if x == y || y == Null {
		return r
	}
	return findPath(x, parent[y], parent, r)
}
