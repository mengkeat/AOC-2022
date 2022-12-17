package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Pt struct{ x, y int }
type Chamber struct {
	solids  map[Pt]bool
	highest int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (c Chamber) Display() {
	for y := c.highest; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			_, ok := c.solids[Pt{x, y}]
			if !ok {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("|")
	}
}

func (c Chamber) Snapshot() [7]int {
	var snap [7]int
	for i := 0; i < 7; i++ {
		snap[i] = -math.MaxInt
	}
	for k, _ := range c.solids {
		snap[k.x] = max(snap[k.x], k.y)
	}
	for i := 0; i < 7; i++ {
		snap[i] = c.highest - snap[i]
	}
	return snap
}

func (c Chamber) Collide(r Rock) bool {
	for _, p := range r.comp {
		if _, ok := c.solids[p]; ok || p.y < 0 || p.x < 0 || p.x >= 7 {
			return true
		}
	}
	return false
}

func (c *Chamber) RestRock(r Rock) {
	for _, p := range r.comp {
		c.solids[p] = true
		if p.y > c.highest {
			c.highest = p.y
		}
	}
}

type Rock struct {
	comp  []Pt
	shape int
}

// Either move <> or down 'v'
func (r Rock) RockMove(c rune) *Rock {
	dx, dy := 1, 0
	if c == '<' {
		dx = -1
	} else if c == 'v' {
		dy = -1
		dx = 0
	}
	new_r := &Rock{
		make([]Pt, 0),
		r.shape}
	for _, p := range r.comp {
		new_r.comp = append(new_r.comp, Pt{p.x + dx, p.y + dy})
	}
	return new_r
}

func newRock(t, sy int) *Rock {
	switch t {
	case 0:
		return &Rock{ // line
			[]Pt{Pt{2, sy}, Pt{3, sy}, Pt{4, sy}, Pt{5, sy}}, 0}
	case 1:
		return &Rock{ // cross
			[]Pt{Pt{3, sy},
				Pt{2, sy + 1}, Pt{3, sy + 1}, Pt{4, sy + 1},
				Pt{3, sy + 2}},
			1}
	case 2:
		return &Rock{ // L
			[]Pt{
				Pt{2, sy}, Pt{3, sy}, Pt{4, sy},
				Pt{4, sy + 1},
				Pt{4, sy + 2}},
			2}
	case 3:
		return &Rock{ // vertical
			[]Pt{Pt{2, sy}, Pt{2, sy + 1}, Pt{2, sy + 2}, Pt{2, sy + 3}},
			3}
	case 4:
		return &Rock{ // square
			[]Pt{Pt{2, sy}, Pt{3, sy},
				Pt{2, sy + 1}, Pt{3, sy + 1}},
			4}
	}
	panic("Not supposed to reach here!")
	return &Rock{[]Pt{}, -1}
}

type SnapKey struct {
	snapshot   [7]int
	last_piece int
	instr_idx  int
}
type SnapVal struct {
	height int
	iter   int
}

func main() {
	bytes, _ := os.ReadFile("input17.txt")
	instr := strings.Split(string(bytes), "\n")[0]

	snap_history := make(map[SnapKey]SnapVal)
	period_len, period_height, offset := -1, -1, -1
	heights := []int{}
	c := Chamber{make(map[Pt]bool), -1}
	icount := 0
	for pcount := 0; pcount < 2022; pcount++ {
		var cur_rock *Rock = newRock(pcount%5, c.highest+4)
		settled := false
		for !settled {
			var lateral_rock *Rock = cur_rock.RockMove(rune(instr[icount]))
			if c.Collide(*lateral_rock) {
				lateral_rock = cur_rock
			}
			var down_rock *Rock = lateral_rock.RockMove('v')
			if c.Collide(*down_rock) {
				c.RestRock(*lateral_rock)
				settled = true
			} else {
				cur_rock = down_rock
			}
			icount = (icount + 1) % len(instr)
		}
		heights = append(heights, c.highest+1)

		skey := SnapKey{c.Snapshot(), (pcount + 1) % 5, icount}
		if _, ok := snap_history[skey]; !ok {
			snap_history[skey] = SnapVal{c.highest + 1, pcount}
		} else {
			period_height = c.highest + 1 - snap_history[skey].height
			period_len = pcount - snap_history[skey].iter
			snap_history[skey] = SnapVal{c.highest + 1, pcount}
			offset = snap_history[skey].iter
		}
	}
	fmt.Println("Part A:", heights[2021])
	fmt.Println(period_len, period_height, offset)

	num_cycles := (1000000000000 - 2022) / period_len
	rem := (1000000000000 - 2022) % period_len
	tgt := heights[2021] + num_cycles*period_height + heights[offset-period_len+rem] - heights[offset-period_len]
	fmt.Println("Part B:", tgt)
}
