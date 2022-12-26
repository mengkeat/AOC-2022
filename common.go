package main

import (
	"os"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

var DirX = []int{1, 0, -1, 0}
var DirY = []int{0, 1, 0, -1}

var Neigh4 = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var Neigh8 = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

type Pt3D struct{ x, y, z int }
type Pt2D struct{ x, y int }

func (p Pt2D) Adj(neigh [][]int) []Pt2D {
	ans := make([]Pt2D, len(neigh))
	for i, d := range neigh {
		ans[i] = Pt2D{p.x + d[0], p.y + d[1]}
	}
	return ans
}
func (p Pt2D) Adj4() []Pt2D { return p.Adj(Neigh4) }
func (p Pt2D) Adj8() []Pt2D { return p.Adj(Neigh8) }

func InMap[T comparable, U any](p T, m map[T]U) bool {
	_, ok := m[p]
	return ok
}

func LinesFromFile(f string) []string {
	bytes, _ := os.ReadFile(f)
	return strings.Split(string(bytes), "\n")
}

func Map[T1, T2 any](source []T1, mapper func(T1) T2) []T2 {
	answer := make([]T2, len(source))
	for i, v := range source {
		answer[i] = mapper(v)
	}
	return answer
}

func Filter[T any](source []T, filter func(T) bool) []T {
	var answer []T
	for _, v := range source {
		if filter(v) {
			answer = append(answer, v)
		}
	}
	return answer
}

func Reduce[T1, T2 any](source []T1, initialValue T2, reducer func(T2, T1) T2) T2 {
	r := initialValue
	for _, v := range source {
		r = reducer(r, v)
	}

	return r
}
