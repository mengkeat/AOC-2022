package main

import (
	"fmt"
	"os"
	"strings"
)

type Pt struct{ r, c int }
type PtDist struct {
	Pt
	dst int
}

var dr []int = []int{0, 1, 0, -1}
var dc []int = []int{1, 0, -1, 0}

func dfs(Q []PtDist, hill [][]int, endpt Pt) int {
	S := make(map[Pt]int)
	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]
		if _, ok := S[curr.Pt]; ok {
			continue
		}
		S[curr.Pt] = 1
		if curr.Pt == endpt {
			return curr.dst
		}
		for d := 0; d < 4; d++ {
			r2, c2 := curr.r+dr[d], curr.c+dc[d]
			if r2 < 0 || r2 >= len(hill) || c2 < 0 || c2 >= len(hill[0]) {
				continue
			}
			if hill[curr.r][curr.c]+1 >= hill[r2][c2] {
				Q = append(Q, PtDist{Pt{r2, c2}, curr.dst + 1})
			}
		}
	}
	return -1
}

func main() {
	bytes, _ := os.ReadFile("input12.txt")
	lines := strings.Split(string(bytes), "\n")

	Q := make([]PtDist, 0)
	Q2 := make([]PtDist, 0)

	endpt := Pt{}
	hill := make([][]int, len(lines))
	for r, ln := range lines {
		hill[r] = make([]int, len(lines[0]))
		for c, ch := range ln {
			switch ch {
			case 'S':
				Q = append(Q, PtDist{Pt{r, c}, 0})
				Q2 = append(Q2, PtDist{Pt{r, c}, 0})
				hill[r][c] = 0
			case 'E':
				hill[r][c] = 25
				endpt = Pt{r, c}
			case 'a':
				Q2 = append(Q2, PtDist{Pt{r, c}, 0})
				hill[r][c] = 0
			default:
				hill[r][c] = int(ch - 'a')
			}
		}
	}
	fmt.Println("Part A:", dfs(Q, hill, endpt))
	fmt.Println("Part B:", dfs(Q2, hill, endpt))
}
