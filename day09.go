package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dir [4][2]int = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Point struct{ x, y int }

func (p Point) DistSq(p2 Point) int {
	dx, dy := p.x-p2.x, p.y-p2.y
	return dx*dx + dy*dy
}

type Rope struct {
	heads, tails Point
	tail_map     map[Point]int
}

func newRope() *Rope {
	r := &Rope{Point{}, Point{}, make(map[Point]int)}
	r.tail_map[Point{}] = 1
	return r
}

func (r *Rope) CorrectTail() {
	best_dist := 20
	ht_dist := r.heads.DistSq(r.tails)
	if ht_dist == 8 {
		r.tails.x = (r.tails.x + r.heads.x) / 2
		r.tails.y = (r.tails.y + r.heads.y) / 2
	} else if ht_dist > 2 {
		best_tail := Point{}
		for d2 := 0; d2 < 4; d2++ {
			curr := Point{r.heads.x + dir[d2][0], r.heads.y + dir[d2][1]}
			curr_dist := r.tails.DistSq(curr)
			if curr_dist < best_dist {
				best_dist = curr_dist
				best_tail = curr
			}
		}
		r.tails = best_tail
	}
	r.tail_map[r.tails] = 1
}

func (r *Rope) Move(d int, steps int) {
	for s := 0; s < steps; s++ {
		r.heads.x += dir[d][0]
		r.heads.y += dir[d][1]
		r.CorrectTail()
	}
}

func MoveLongRope(long_rope []*Rope, d int, steps int) {
	for s := 0; s < steps; s++ {
		long_rope[0].Move(d, 1)
		for i := 1; i < len(long_rope); i++ {
			long_rope[i].heads = long_rope[i-1].tails
			long_rope[i].CorrectTail()
		}
	}
}

func main() {
	bytes, _ := os.ReadFile("input09.txt")
	lines := strings.Split(string(bytes), "\n")

	rope := newRope()
	var long_rope []*Rope = make([]*Rope, 0)
	for i := 0; i < 10; i++ {
		long_rope = append(long_rope, newRope())
	}

	for _, instr := range lines {
		tok := strings.Split(instr, " ")
		steps, _ := strconv.Atoi(tok[1])
		switch tok[0] {
		case "R":
			rope.Move(0, steps)
			MoveLongRope(long_rope, 0, steps)
		case "U":
			rope.Move(1, steps)
			MoveLongRope(long_rope, 1, steps)
		case "L":
			rope.Move(2, steps)
			MoveLongRope(long_rope, 2, steps)
		case "D":
			rope.Move(3, steps)
			MoveLongRope(long_rope, 3, steps)
		}
	}
	fmt.Println("Part A:", len(rope.tail_map))
	fmt.Println("Part B:", len(long_rope[8].tail_map))
}
