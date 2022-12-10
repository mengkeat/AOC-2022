package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetStrength(s map[int]int, cycle int) int {
	v, ok := s[cycle]
	if ok {
		return v
	} else {
		return s[cycle-1]
	}
}

func main() {
	bytes, _ := os.ReadFile("input10.txt")
	lines := strings.Split(string(bytes), "\n")

	cycle := 1
	X := 1
	smap := map[int]int{1: 1}
	for _, ln := range lines {
		tok := strings.Split(ln, " ")
		switch tok[0] {
		case "noop":
			cycle += 1
		case "addx":
			v, _ := strconv.Atoi(tok[1])
			X += v
			cycle += 2
		}
		smap[cycle] = X
	}
	part_a := 0
	for s := 20; s <= 220; s += 40 {
		strength := GetStrength(smap, s)
		part_a += s * strength
	}
	fmt.Println("Part A:", part_a)

	crt := make([][]rune, 6)
	for c := 0; c < 6*40; c++ {
		curr_row := c / 40
		if c%40 == 0 {
			crt[curr_row] = make([]rune, 40)
		}
		sprite_pos := GetStrength(smap, c+1)
		curr_col := c % 40
		if curr_col >= sprite_pos-1 && curr_col <= sprite_pos+1 {
			crt[curr_row][curr_col] = rune('#')
		} else {
			crt[curr_row][curr_col] = rune('.')
		}
	}
	for _, row := range crt {
		fmt.Println(string(row))
	}
}
