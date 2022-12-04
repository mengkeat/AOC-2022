package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input04.txt")
	lines := strings.Split(string(bytes), "\n")

	count, count2 := 0, 0
	for _, s := range lines {
		tok := strings.Split(strings.ReplaceAll(s, "-", ","), ",")
		a1, _ := strconv.Atoi(tok[0])
		a2, _ := strconv.Atoi(tok[1])
		b1, _ := strconv.Atoi(tok[2])
		b2, _ := strconv.Atoi(tok[3])

		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			count++
		}
		if !(a2 < b1 || a1 > b2) {
			count2++
		}
	}
	fmt.Println("Part A:", count)
	fmt.Println("Part B:", count2)
}
