package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Eqn struct {
	left, right string
	op          string
	val         int
	has_humn    bool
}

func compute(all map[string]*Eqn, s string) int {
	left, right := 0, 0
	if len(all[s].left) == 0 && len(all[s].right) == 0 {
		return all[s].val
	}

	left = compute(all, all[s].left)
	right = compute(all, all[s].right)
	all[s].has_humn = all[all[s].left].has_humn || all[all[s].right].has_humn

	switch all[s].op {
	case "+":
		all[s].val = left + right
	case "-":
		all[s].val = left - right
	case "*":
		all[s].val = left * right
	case "/":
		all[s].val = left / right
	}
	return all[s].val
}

func find(m string, tgt int, all map[string]*Eqn) (string, int) {
	cur := all[m]
	switch cur.op {
	case "+":
		if all[cur.left].has_humn {
			return cur.left, tgt - all[cur.right].val
		} else {
			return cur.right, tgt - all[cur.left].val
		}
	case "-":
		if all[cur.left].has_humn {
			return cur.left, tgt + all[cur.right].val
		} else {
			return cur.right, all[cur.left].val - tgt
		}
	case "*":
		if all[cur.left].has_humn {
			return cur.left, tgt / all[cur.right].val
		} else {
			return cur.right, tgt / all[cur.left].val
		}
	case "/":
		if all[cur.left].has_humn {
			return cur.left, tgt * all[cur.right].val
		} else {
			return cur.right, all[cur.left].val / tgt
		}
	}
	return "", 0
}

func resolv(m string, tgt int, all map[string]*Eqn) int {
	var nxt_tgt int = tgt
	for m != "humn" {
		m, nxt_tgt = find(m, nxt_tgt, all)
	}
	return nxt_tgt
}

func main() {
	lines := LinesFromFile("input21.txt")

	all := make(map[string]*Eqn)
	for _, ln := range lines {
		tok := strings.Split(ln, " ")
		if len(tok) < 3 {
			x, _ := strconv.Atoi(tok[1])
			all[tok[0][:4]] = &Eqn{val: x}
		} else {
			all[tok[0][:4]] = &Eqn{tok[1], tok[3], tok[2], -1, false}
		}
	}
	all["humn"].has_humn = true
	part_a := compute(all, "root")
	fmt.Println("Part A:", part_a)

	all["root"].op = "-"
	part_b := resolv("root", 0, all)
	fmt.Println("Part B:", part_b)
}
