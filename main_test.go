package main

import "testing"

func TestHasCycle(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, true},
		{3, [][2]int{{0, 1}, {1, 2}}, false},
		{4, [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, false}, // diamond DAG
		{1, [][2]int{{0, 0}}, true},                          // self loop
		{5, [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true},  // cycle beside a path
		{2, nil, false},
	}
	for _, c := range cases {
		if got := HasCycle(c.n, c.edges); got != c.want {
			t.Errorf("HasCycle(n=%d, %v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
