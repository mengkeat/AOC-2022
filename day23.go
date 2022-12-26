package main

import "fmt"

type ElvesMap = map[Pt2D]int // 2D pos -> decision

func MoveElves(m1 ElvesMap, round int) (ElvesMap, bool) {
	new_map := make(ElvesMap)
	occupied := make(map[Pt2D][]Pt2D) // Elves that want to occupy

	noOne := func(p Pt2D, neigh [][]int) bool {
		alone := true
		for _, d := range neigh {
			alone = alone && !InMap(Pt2D{p.x + d[0], p.y + d[1]}, m1)
		}
		return alone
	}
	isAlone := func(p Pt2D) bool { return noOne(p, Neigh8) }
	northOK := func(p Pt2D) bool { return noOne(p, [][]int{{-1, -1}, {-1, 0}, {-1, 1}}) }
	southOK := func(p Pt2D) bool { return noOne(p, [][]int{{1, -1}, {1, 0}, {1, 1}}) }
	westOK := func(p Pt2D) bool { return noOne(p, [][]int{{-1, -1}, {0, -1}, {1, -1}}) }
	eastOK := func(p Pt2D) bool { return noOne(p, [][]int{{-1, 1}, {0, 1}, {1, 1}}) }
	dir_decision := []func(Pt2D) bool{northOK, southOK, westOK, eastOK}
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	no_move_count := 0
	for pos, dec := range m1 {
		pos_to_add := pos
		if isAlone(pos) {
			pos_to_add = pos
			new_map[pos_to_add] = dec + 1
			no_move_count++
		} else {
			moved := false
			for st := 0; st < 4; st++ {
				i := ((round % 4) + st) % 4
				if dir_decision[i](pos) {
					pos_to_add = Pt2D{pos.x + dir[i][0], pos.y + dir[i][1]}
					new_map[pos_to_add] = dec + 1
					moved = true
					break
				}
			}
			if moved == false {
				pos_to_add = pos
				new_map[pos_to_add] = dec + 1
				no_move_count++
			}
		}
		if !InMap(pos_to_add, occupied) {
			occupied[pos_to_add] = make([]Pt2D, 0)
		}
		occupied[pos_to_add] = append(occupied[pos_to_add], pos)
	}

	// Move overlapping ones back
	for p, v := range occupied {
		if len(v) > 1 {
			delete(new_map, p)
			for _, pos := range v {
				new_map[pos] = round
			}
		}
	}

	return new_map, (no_move_count == len(m1))
}

func GetEmpty(m ElvesMap) int {
	xs, ys := []int{}, []int{}
	for p, _ := range m {
		xs = append(xs, p.x)
		ys = append(ys, p.y)
	}
	x_range := Max(xs...) - Min(xs...) + 1
	y_range := Max(ys...) - Min(ys...) + 1
	return x_range*y_range - len(m)
}

func main() {
	lines := LinesFromFile("input23.txt")

	elves := make(ElvesMap) // 2D position -> first choice dir
	for r, ln := range lines {
		for c, ch := range ln {
			if ch == '#' {
				elves[Pt2D{r, c}] = 0
			}
		}
	}

	for i, no_moves := 0, false; no_moves == false; i++ {
		elves, no_moves = MoveElves(elves, i)
		if i+1 == 10 {
			fmt.Println("Part A:", GetEmpty(elves))
		}
		if no_moves {
			fmt.Println("Part B:", i+1)
		}
	}
}
