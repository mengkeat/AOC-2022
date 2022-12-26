package main

import (
	"fmt"
	"strings"
)

type VWind struct {
	pos           map[Pt2D]rune
	width, height int
	winds_cache   map[int]map[Pt2D]rune
}

func (w VWind) WindsAt(t int) map[Pt2D]rune {
	if InMap(t, w.winds_cache) {
		return w.winds_cache[t]
	}
	new_w := make(map[Pt2D]rune)
	for k, v := range w.pos {
		switch v {
		case '>':
			new_w[Pt2D{k.x, Mod(k.y-1+t, w.width-2) + 1}] = '>'
		case '<':
			new_w[Pt2D{k.x, Mod(k.y-1-t, w.width-2) + 1}] = '<'
		case '^':
			new_w[Pt2D{Mod(k.x-1-t, w.height-2) + 1, k.y}] = '^'
		case 'v':
			new_w[Pt2D{Mod(k.x-1+t, w.height-2) + 1, k.y}] = 'v'
		}
	}

	w.winds_cache[t] = new_w
	return w.winds_cache[t]
}

type VState struct {
	Pt2D
	time int
}

func bfs(start VState, dest Pt2D, winds VWind) int {
	Q := []VState{start}
	visited := make(map[VState]bool)

	Neigh4 = append(Neigh4, []int{0, 0})
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur.Pt2D == dest {
			return cur.time
		}

		if InMap(cur, visited) {
			continue
		}
		visited[cur] = true
		for _, d := range Neigh4 {
			npos := Pt2D{cur.Pt2D.x + d[0], cur.Pt2D.y + d[1]}
			if InMap(VState{npos, cur.time + 1}, visited) {
				continue
			}
			if npos.x <= 0 || npos.x >= winds.height-1 || npos.y <= 0 || npos.y >= winds.width-1 {
				if npos != dest && npos != start.Pt2D {
					continue
				}
			}
			if InMap(npos, winds.WindsAt(cur.time+1)) {
				continue
			}
			Q = append(Q, VState{npos, cur.time + 1})
		}
	}
	return -1
}

func ShowWinds(w VWind, t int) {
	fmt.Println()
	for r := 0; r < w.height; r++ {
		for c := 0; c < w.width; c++ {
			if InMap(Pt2D{r, c}, w.WindsAt(t)) {
				fmt.Print(string(rune(w.WindsAt(t)[Pt2D{r, c}])))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	lines := LinesFromFile("input24.txt")
	start_r, start_c := 0, strings.Index(lines[0], ".")
	end_r, end_c := len(lines)-1, strings.Index(lines[len(lines)-1], ".")

	winds := VWind{make(map[Pt2D]rune), len(lines[0]), len(lines), make(map[int]map[Pt2D]rune)}
	for r, ln := range lines {
		for c, ch := range ln {
			if ch == '>' || ch == '<' || ch == '^' || ch == 'v' {
				winds.pos[Pt2D{r, c}] = ch
			}
		}
	}
	part_a := bfs(VState{Pt2D{start_r, start_c}, 0}, Pt2D{end_r, end_c}, winds)
	fmt.Println("Part A:", part_a)

	back := bfs(VState{Pt2D{end_r, end_c}, part_a}, Pt2D{start_r, start_c}, winds)
	part_b := bfs(VState{Pt2D{start_r, start_c}, back}, Pt2D{end_r, end_c}, winds)
	fmt.Println("Part B:", part_b)
}
