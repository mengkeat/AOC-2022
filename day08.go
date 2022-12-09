package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input08.txt")
	lines := strings.Split(string(bytes), "\n")
	orig := [][]int{}

	for _, ln := range lines {
		row := make([]int, len(ln))
		for i, c := range ln {
			row[i] = int(c) - '0'
		}
		orig = append(orig, row)
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	clamp := func(x, lo, hi int) int {
		if x < lo {
			return lo
		} else if x > hi {
			return hi
		} else {
			return x
		}
	}

	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	nrows, ncols := len(orig), len(orig[0])
	max_scene := 0
	viz_count := 0

	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			vis := false
			scene_val := 1
			for d := 0; d < 4; d++ {
				r2, c2 := r+dir[d][0], c+dir[d][1]
				for r2 >= 0 && r2 < nrows && c2 >= 0 && c2 < ncols && orig[r2][c2] < orig[r][c] {
					r2 += dir[d][0]
					c2 += dir[d][1]
				}
				if r2 < 0 || r2 == nrows || c2 < 0 || c2 == ncols {
					vis = vis || true
				}
				r2, c2 = clamp(r2, 0, nrows-1), clamp(c2, 0, ncols-1)
				scene_val = scene_val * (abs(r2-r) + abs(c2-c))
			}
			if vis {
				viz_count++
			}
			if scene_val > max_scene {
				max_scene = scene_val
			}
		}
	}
	fmt.Println("Part A:", viz_count)
	fmt.Println("Part B:", max_scene)
}
