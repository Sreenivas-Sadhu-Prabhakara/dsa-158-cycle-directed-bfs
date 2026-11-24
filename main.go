package main

import "fmt"

// HasCycle reports whether a directed graph contains a cycle using Kahn's
// BFS: peel vertices of indegree zero; if the peeling cannot consume every
// vertex, the leftovers form (or feed) a cycle.
func HasCycle(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	indegree := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		indegree[e[1]]++
	}
	queue := []int{}
	for v := 0; v < n; v++ {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	processed := 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		processed++
		for _, w := range adj[v] {
			indegree[w]--
			if indegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return processed < n
}

func main() {
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}))
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {1, 2}}))
}
