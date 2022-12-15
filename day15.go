package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Pt struct{ x, y int }
type Zone struct {
	sensor_dist, beacons map[Pt]int
}

func (z Zone) NotBeacon(x, y int) bool {
	for p, d := range z.sensor_dist {
		_, is_beacon := z.beacons[Pt{x, y}]
		if dist(Pt{x, y}, p) <= d && !is_beacon {
			return true
		}
	}
	return false
}

func dist(a, b Pt) int {
	dx, dy := a.x-b.x, a.y-b.y
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

func main() {

	bytes, _ := os.ReadFile("input15.txt")
	// bytes, _ := os.ReadFile("test15.txt")
	lines := strings.Split(string(bytes), "\n")

	z := Zone{make(map[Pt]int), make(map[Pt]int)}
	maxdist := -1
	minx, maxx := math.MaxInt, math.MinInt
	for _, ln := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(ln, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&x1, &y1, &x2, &y2)
		z.sensor_dist[Pt{x1, y1}] = dist(Pt{x1, y1}, Pt{x2, y2})
		z.beacons[Pt{x2, y2}] = 1
		if minx > x1 {
			minx = x1
		}
		if minx > x2 {
			minx = x2
		}
		if maxx < x1 {
			maxx = x1
		}
		if maxx < x2 {
			maxx = x2
		}
		if maxdist < z.sensor_dist[Pt{x1, y1}] {
			maxdist = z.sensor_dist[Pt{x1, y1}]
		}
	}
	y := 2000000
	part_a := 0
	for x := minx - maxdist; x <= maxx+maxdist; x++ {
		if z.NotBeacon(x, y) {
			part_a++
		}
	}
	fmt.Println("Part A:", part_a)

	max := 4000000
	part_b := Pt{}
	for s, d := range z.sensor_dist {
		for x, i := s.x-d-1, 0; x <= s.x; x, i = x+1, i+1 {
			if x < 0 || x > max {
				continue
			}
			if !z.NotBeacon(x, s.y-i) {
				if s.y-i < 0 || s.y-i > max {
					continue
				}
				part_b = Pt{x, s.y - i}
			}
			if !z.NotBeacon(x, s.y+i) {
				if s.y+i < 0 || s.y+i > max {
					continue
				}
				part_b = Pt{x, s.y + i}
			}
		}
		for x, i := s.x, d+1; x <= s.x+d+1; x, i = x+1, i-1 {
			if x < 0 || x > max {
				continue
			}
			if !z.NotBeacon(x, s.y-i) {
				if s.y-i < 0 || s.y-i > max {
					continue
				}
				part_b = Pt{x, s.y - i}
			}
			if !z.NotBeacon(x, s.y+i) {
				if s.y+i < 0 || s.y+i > max {
					continue
				}
				part_b = Pt{x, s.y + i}
			}
		}
	}
	fmt.Println("Part B:", part_b.x*4000000+part_b.y)
}
