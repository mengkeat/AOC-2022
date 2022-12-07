package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input07.txt")
	lines := strings.Split(string(bytes), "\n")

	sz_count := map[string]int{}
	path := []string{}
	for _, ln := range lines {
		tok := strings.Split(ln, " ")
		if tok[1] == "cd" {
			if tok[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, tok[2])
			}
		} else if tok[0] == "dir" || tok[1] == "ls" {
			continue
		} else {
			sz, _ := strconv.Atoi(tok[0])
			for i := 1; i < len(path)+1; i++ {
				full_path := strings.Join(path[:i], "/")
				sz_count[full_path] += sz
			}
		}
	}
	part_a, part_b := 0, sz_count["/"]+1
	free_criteria := sz_count["/"] - 40000000
	for _, v := range sz_count {
		if v < 100000 {
			part_a += v
		}
		if v >= free_criteria {
			if v < part_b {
				part_b = v
			}
		}
	}
	fmt.Println("Part A:", part_a)
	fmt.Println("Part B:", part_b)
}
