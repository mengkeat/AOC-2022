package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input02.txt")
	lines := strings.Split(string(bytes), "\n")

	score, score2 := 0, 0
	for _, s := range lines {
		if s == "" {
			continue
		}
		opp := s[0] - 'A'
		me := s[2] - 'X'

		if (me+1)%3 == opp { // Lost
			score += int(me) + 1
		} else if me == opp { // Draw
			score += 3 + int(me) + 1
		} else {
			score += 6 + int(me) + 1
		}

		if me == 0 { // Lost
			score2 += (int(opp)+2)%3 + 1
		} else if me == 1 { // draw
			score2 += 3 + int(opp) + 1
		} else {
			score2 += 6 + (int(opp)+1)%3 + 1
		}
	}
	fmt.Println("Part A:", score)
	fmt.Println("Part B:", score2)
}
