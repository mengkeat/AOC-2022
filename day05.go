package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct{ arr []rune }

func (s *Stack) push(r []rune) {
	s.arr = append(s.arr, r...)
}

func (s *Stack) pop_n(n int) {
	s.arr = s.arr[:len(s.arr)-n]
}

func (s *Stack) reverse_top_n(n int) {
	for i, j := len(s.arr)-n, len(s.arr)-1; i < j; i, j = i+1, j-1 {
		s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
	}
}

func (s Stack) top() rune {
	return s.arr[len(s.arr)-1]
}

func (s Stack) top_n(n int) []rune {
	return s.arr[len(s.arr)-n:]
}

func main() {
	file, _ := os.Open("input05.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	st := make([]Stack, 9)
	st2 := make([]Stack, 9)

	sc.Scan()
	for sc.Text()[1] != '1' {
		for i, c := range sc.Text() {
			if c >= 'A' && c <= 'Z' {
				st[i/4].arr = append([]rune{c}, st[i/4].arr...)
				st2[i/4].arr = append([]rune{c}, st2[i/4].arr...)
			}
		}
		sc.Scan()
	}
	sc.Scan()

	for sc.Scan() {
		sz, from, to := -1, -1, -1
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &sz, &from, &to)
		st[to-1].push(st[from-1].top_n(sz))
		st[to-1].reverse_top_n(sz)
		st[from-1].pop_n(sz)

		st2[to-1].push(st2[from-1].top_n(sz))
		st2[from-1].pop_n(sz)
	}

	fmt.Print("Part A: ")
	for _, s := range st {
		fmt.Print(string(s.top()))
	}

	fmt.Print("\nPart B: ")
	for _, s := range st2 {
		fmt.Print(string(s.top()))
	}
}
