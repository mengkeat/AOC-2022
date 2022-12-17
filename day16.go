package main

import (
	"fmt"
	"os"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Key struct {
	cur, visited, release, time_left int
}
type Cave struct {
	flow    map[int]int
	dist    [][]int
	n_nodes int
	start   int
	DP      map[Key]int
	DP_2    map[Key]int
}

func (c Cave) dfs_alone(cur int, visited int, release int, time_left int) int {
	cur_key := Key{cur, visited, release, time_left}
	if _, ok := c.DP[cur_key]; ok {
		return c.DP[cur_key]
	}

	if time_left <= 0 {
		return release
	}
	local_v := release
	if c.flow[cur] > 0 {
		time_left = time_left - 1
		local_v += time_left * c.flow[cur]
	}
	for i := 0; i < c.n_nodes; i++ {
		next_visited := (1 << i) | visited
		if next_visited != visited && time_left > c.dist[cur][i] && c.flow[i] > 0 {
			local_v = max(local_v,
				c.dfs_alone(i, next_visited,
					release+(time_left*c.flow[cur]),
					time_left-c.dist[cur][i]))
		}
	}
	c.DP[cur_key] = local_v
	return local_v
}

func (c Cave) dfs_elephant(cur int, visited int, release int, time_left int) int {
	cur_key := Key{cur, visited, release, time_left}
	if _, ok := c.DP_2[cur_key]; ok {
		return c.DP_2[cur_key]
	}

	if time_left <= 0 {
		return release
	}
	local_v := release
	if c.flow[cur] > 0 {
		time_left = time_left - 1
		local_v += time_left * c.flow[cur]
	}
	for i := 0; i < c.n_nodes; i++ {
		next_visited := (1 << i) | visited
		if next_visited != visited && time_left > c.dist[cur][i] && c.flow[i] > 0 {
			local_v = max(local_v,
				c.dfs_elephant(i, next_visited,
					release+(time_left*c.flow[cur]),
					time_left-c.dist[cur][i]))
		}
		local_v = max(local_v, c.dfs_alone(c.start, visited, release+(time_left*c.flow[cur]), 26))
	}
	c.DP_2[cur_key] = local_v
	return local_v
}

func main() {
	bytes, _ := os.ReadFile("input16.txt")
	// bytes, _ := os.ReadFile("test16.txt")
	lines := strings.Split(string(bytes), "\n")
	cave := Cave{make(map[int]int), make([][]int, len(lines)), 0, 0,
		make(map[Key]int), make(map[Key]int)}
	for i := 0; i < len(lines); i++ {
		cave.dist[i] = make([]int, len(lines))
		for j := 0; j < len(lines); j++ {
			cave.dist[i][j] = 100000
		}
	}

	// Golang can get quite verboose
	count := 0
	num_map := make(map[string]int)
	node_map := func(s string) int {
		if _, ok := num_map[s]; !ok {
			num_map[s] = count
			count++
		}
		return num_map[s]
	}

	var node_str string
	var rate int
	for _, ln := range lines {
		fmt.Sscanf(ln, "Valve %s has flow rate=%d;", &node_str, &rate)
		f := strings.Fields(strings.ReplaceAll(ln, ",", ""))
		n := node_map(node_str)
		cave.flow[n] = rate
		for _, adj := range f[9:] {
			adj_n := node_map(adj)
			cave.dist[n][adj_n] = 1
		}
	}

	cave.n_nodes = len(lines)
	for k := 0; k < cave.n_nodes; k++ {
		for i := 0; i < cave.n_nodes; i++ {
			for j := 0; j < cave.n_nodes; j++ {
				cave.dist[i][j] = min(cave.dist[i][j],
					cave.dist[i][k]+cave.dist[k][j])
			}
		}
	}

	cave.start = node_map("AA")
	rel_a := cave.dfs_alone(cave.start, 0, 0, 30)
	fmt.Println("Part A:", rel_a)

	rel_b := cave.dfs_elephant(cave.start, 0, 0, 26)
	fmt.Println("Part B:", rel_b)
}
