package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pt struct{ r, c int }
type Reservoir struct {
	grid     map[Pt]rune
	maxdepth int
	numrocks int
}

func (r *Reservoir) FillRocks(s string) {
	s = strings.Replace(s, " -> ", ",", -1)
	tok := strings.Split(s, ",")
	c := make([]int, len(tok))
	for i, x := range tok {
		c[i], _ = strconv.Atoi(x)
	}
	for i := 0; i < len(c)-2; i = i + 2 {
		x1, y1, x2, y2 := c[i], c[i+1], c[i+2], c[i+3]
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if y2 > r.maxdepth {
			r.maxdepth = y2
		}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				r.grid[Pt{y, x}] = '#'
			}
		}
	}
}

func (r *Reservoir) Pour(curr Pt, part_a bool) bool {
	_, ok := r.grid[curr]
	if ok || (!part_a && curr.r >= r.maxdepth+2) {
		return false
	}
	if part_a && curr.r > r.maxdepth {
		return true
	}
	for _, d := range [][]int{{0, 1}, {-1, 1}, {1, 1}} {
		nxt := Pt{curr.r + d[1], curr.c + d[0]}
		if r.Pour(nxt, part_a) {
			return true
		}
	}
	r.grid[curr] = 'O'
	return false
}

func main() {
	bytes, _ := os.ReadFile("input14.txt")
	lines := strings.Split(string(bytes), "\n")

	res := Reservoir{make(map[Pt]rune), -1, -1}
	for _, ln := range lines {
		res.FillRocks(ln)
	}
	res.numrocks = len(res.grid)

	res.Pour(Pt{0, 500}, true)
	fmt.Println("Part A:", len(res.grid)-res.numrocks)

	res.Pour(Pt{0, 500}, false)
	fmt.Println("Part B:", len(res.grid)-res.numrocks)
}
