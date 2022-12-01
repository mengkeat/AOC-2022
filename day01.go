package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input01.txt")
	lines := strings.Split(string(bytes), "\n")

	var sums []int
	curr := 0
	max := -1
	for _, s := range lines {
		if s == "" {
			if curr > max {
				max = curr
			}
			sums = append(sums, curr)
			curr = 0
		} else {
			i, _ := strconv.Atoi(s)
			curr += i
		}
	}
	fmt.Println("Part A:", max)

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println("Part B:", sums[0]+sums[1]+sums[2])
}
