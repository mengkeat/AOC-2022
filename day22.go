package main

import (
	"fmt"
)

var L int = 50

type Move struct {
	turn  rune
	steps int
}

type Tiles struct {
	width, height int
	m             [][]rune
	start_col     []int
	end_col       []int
}

func (t *Tiles) ReadData(s string) []Move {
	lines := LinesFromFile(s)

	t.width, t.height = 0, 0
	for t.height = 0; len(lines[t.height]) != 0; t.height++ {
		t.width = Max(t.width, len(lines[t.height]))
	}

	t.start_col, t.end_col = make([]int, t.height), make([]int, t.height)
	t.m = make([][]rune, t.height)
	for i := 0; i < t.height; i++ {
		t.m[i] = make([]rune, t.width)
		a, b := t.width*2, -1
		for j := 0; j < t.width; j++ {
			if j >= len(lines[i]) || lines[i][j] == ' ' {
				t.m[i][j] = rune(' ')
			} else {
				a = Min(a, j)
				b = Max(b, j)
				t.m[i][j] = rune(lines[i][j])
			}
		}
		t.start_col[i], t.end_col[i] = a, b
	}

	moves := []Move{}
	cur := 0
	for _, c := range lines[t.height+1] {
		if c >= '0' && c <= '9' {
			cur = cur*10 + int(c-'0')
		} else {
			moves = append(moves, Move{'X', cur}, Move{c, 0})
			cur = 0
		}
	}
	if cur > 0 {
		moves = append(moves, Move{'X', cur})
	}
	return moves
}

// r, c is already at position where its in a blank space
func (t Tiles) GotoEdge(r, c, dr, dc int) (int, int) {
	or, oc := r-dr, c-dc
	for t.m[r][c] == ' ' {
		r, c = (r+dr+t.height)%t.height, (c+dc+t.width)%t.width
	}
	if t.m[r][c] == '#' {
		return oc, or
	}
	return c, r
}

// Transpose will be performed first before a flip
func transform(r, c, d, new_rr, new_cc int, flip_r, flip_c, transpose bool) (int, int, int) {
	rr_mod, cc_mod := r%L, c%L
	new_r, new_c, new_d := rr_mod, cc_mod, d
	if transpose {
		new_r, new_c = new_c, new_r
		if new_d == 0 {
			new_d = 1
		} else if new_d == 1 {
			new_d = 0
		} else if new_d == 2 {
			new_d = 3
		} else if new_d == 3 {
			new_d = 2
		}
	}
	if flip_r {
		new_r = L - new_r - 1
		if new_d == 1 {
			new_d = 3
		} else if new_d == 3 {
			new_d = 1
		}
	}
	if flip_c {
		new_c = L - new_c - 1
		if new_d == 0 {
			new_d = 2
		} else if new_d == 2 {
			new_d = 0
		}
	}
	new_r += new_rr * L
	new_c += new_cc * L
	return new_r, new_c, new_d
}

func CubeWrap(r, c, d int) (int, int, int) {
	RR := r / L
	CC := c / L
	new_r, new_c, new_d := 0, 0, 0
	if RR == 0 && CC == 0 {
		switch d {
		case 0:
			new_r, new_c, new_d = transform(r, c, d, 2, 1, true, true, false)
		case 1:
			new_r, new_c, new_d = transform(r, c, d, 0, 2, false, false, false)
		case 2:
			new_r, new_c, new_d = transform(r, c, d, 2, 0, true, true, false)
		case 3:
			panic("Impossible direction 3 for Region (0,0)")
		}
	} else if RR == 1 && CC == 0 {
		switch d {
		case 0, 1:
			panic("Impossible direction 0/1 for Region (1,0)")
		case 2:
			new_r, new_c, new_d = transform(r, c, d, 2, 0, true, false, true)
		case 3:
			new_r, new_c, new_d = transform(r, c, d, 1, 1, false, true, true)
		}

	} else if RR == 1 && CC == 2 {
		switch d {
		case 0:
			new_r, new_c, new_d = transform(r, c, d, 0, 2, true, false, true)
		case 1:
			new_r, new_c, new_d = transform(r, c, d, 1, 1, false, true, true)
		case 2, 3:
			panic("Impossible direction 2/3 for region (1,2)")
		}
	} else if RR == 2 && CC == 2 {
		switch d {
		case 0:
			new_r, new_c, new_d = transform(r, c, d, 0, 2, true, true, false)
		case 2:
			new_r, new_c, new_d = transform(r, c, d, 0, 1, true, true, false)
		case 1, 3:
			panic("Imppossible direction 1/3 for region (2,2)")
		}
	} else if RR == 3 && CC == 1 {
		switch d {
		case 0:
			new_r, new_c, new_d = transform(r, c, d, 2, 1, true, false, true)
		case 1:
			new_r, new_c, new_d = transform(r, c, d, 3, 0, false, true, true)
		case 3:
			new_r, new_c, new_d = transform(r, c, d, 3, 0, false, true, true)
		case 2:
			panic("Impossible direction 2 for region (3, 1)")
		}

	} else if RR == 3 && CC == 2 {
		switch d {
		case 0, 1:
			panic("Impossible direction 0/1 for region (3,2)")
		case 2:
			new_r, new_c, new_d = transform(r, c, d, 0, 1, true, false, true)
		case 3:
			new_r, new_c, new_d = transform(r, c, d, 3, 0, false, false, false)
		}
	}
	return new_r, new_c, new_d
}

func (t Tiles) GetPassword(mv []Move, part int) int {
	d := 0
	c, r := t.start_col[0], 0

	for _, mv := range mv {
		// Move #steps
		for mv.steps > 0 {
			oc, or, od := c, r, d
			c, r = (c+DirX[d]+t.width)%t.width, (r+DirY[d]+t.height)%t.height
			if t.m[r][c] == '#' {
				c, r, d = oc, or, od
				break
			} else if t.m[r][c] == ' ' {
				if part == 1 {
					c, r = t.GotoEdge(r, c, DirY[d], DirX[d])
				} else {
					new_r, new_c, new_d := CubeWrap(r, c, d)
					if t.m[new_r][new_c] == '#' {
						c, r, d = oc, or, od
						break
					} else {
						c, r, d = new_c, new_r, new_d
					}
				}
			}
			mv.steps--
		}
		// Turn
		if mv.turn == 'R' {
			d = (d + 1) % 4
		} else if mv.turn == 'L' {
			d = (d + 3) % 4
		}
	}
	return 1000*(r+1) + 4*(c+1) + d
}

func main() {
	tile := Tiles{}
	moves := tile.ReadData("input22.txt")

	part_a := tile.GetPassword(moves, 1)
	fmt.Println("Part A:", part_a)

	part_b := tile.GetPassword(moves, 2)
	fmt.Println("Part B:", part_b)
}
