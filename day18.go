package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

type Pt3D struct{ x, y, z int }

func inmap(p Pt3D, pts map[Pt3D]bool) bool {
	_, ok := pts[p]
	return ok
}

func GetSurfaceArea(pts map[Pt3D]bool) int {
	ncount := 0
	for p, _ := range pts {
		for _, d := range [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0},
			{0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
			if inmap(Pt3D{p.x + d[0], p.y + d[1], p.z + d[2]}, pts) {
				ncount++
			}
		}
	}
	return len(pts)*6 - ncount
}

func explore(cur Pt3D, lava map[Pt3D]bool, explored map[Pt3D]bool) {
	if len(explored) > 5000 {
		return
	}
	explored[cur] = true
	for _, d := range [6][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0},
		{0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {

		next_pt := Pt3D{cur.x + d[0], cur.y + d[1], cur.z + d[2]}
		if !inmap(next_pt, lava) && !inmap(next_pt, explored) {
			explore(next_pt, lava, explored)
		}
	}
	return
}

func main() {
	bytes, _ := os.ReadFile("input18.txt")
	lines := strings.Split(string(bytes), "\n")

	lava := make(map[Pt3D]bool)
	minx, miny, minz := math.MaxInt, math.MaxInt, math.MaxInt
	maxx, maxy, maxz := math.MinInt, math.MinInt, math.MinInt
	for _, ln := range lines {
		tok := strings.Split(ln, ",")
		x, _ := strconv.Atoi(tok[0])
		y, _ := strconv.Atoi(tok[1])
		z, _ := strconv.Atoi(tok[2])
		lava[Pt3D{x, y, z}] = true
		minx, miny, minz = min(minx, x), min(miny, y), min(minz, z)
		maxx, maxy, maxz = max(maxx, x), max(maxy, y), max(maxz, z)
	}

	lava_surface := GetSurfaceArea(lava)
	fmt.Println("Part A:", lava_surface)

	inside := map[Pt3D]bool{}
	outside := map[Pt3D]bool{}
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			for z := minz; z <= maxz; z++ {
				curr := Pt3D{x, y, z}
				if !inmap(curr, outside) && !inmap(curr, lava) {
					temp := make(map[Pt3D]bool)
					explore(Pt3D{x, y, z}, lava, temp)
					if len(temp) < 5000 {
						for k, v := range temp {
							inside[k] = v
						}
					} else {
						for k, v := range temp {
							outside[k] = v
						}
					}
				}
			}
		}
	}
	air_surf_area := GetSurfaceArea(inside)
	fmt.Println("Part_B:", lava_surface-air_surf_area)
}
