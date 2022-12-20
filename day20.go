package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mix(ele []int, times int) int {
	orig := []*list.Element{}
	lst := list.New()

	for _, e := range ele {
		orig = append(orig, lst.PushBack(e))
	}

	move_forward := func(l *list.List, start *list.Element, n int) *list.Element {
		p := start
		for i := 0; i < n; i++ {
			p = p.Next()
			if p == nil {
				p = l.Front()
			}
		}
		return p
	}

	for t := 0; t < times; t++ {
		for i, p := range orig {
			mv := p.Value.(int)
			s := p
			if mv > 0 {
				s = move_forward(lst, s, mv%(lst.Len()-1))
				lst.MoveAfter(orig[i], s)
			} else {
				new_mv := (-mv) % (lst.Len() - 1)
				for c := 0; c < (new_mv); c++ {
					s = s.Prev()
					if s == nil {
						s = lst.Back()
					}
				}
				lst.MoveBefore(orig[i], s)
			}
		}
	}

	zero := lst.Front()
	for zero.Value.(int) != 0 {
		zero = zero.Next()
	}

	n1 := move_forward(lst, zero, 1000)
	n2 := move_forward(lst, zero, 2000)
	n3 := move_forward(lst, zero, 3000)

	return n1.Value.(int) + n2.Value.(int) + n3.Value.(int)
}

func main() {
	bytes, _ := os.ReadFile("input20.txt")

	ele := []int{}
	for _, i := range strings.Split(string(bytes), "\n") {
		e, _ := strconv.Atoi(i)
		ele = append(ele, e)
	}
	part_a := Mix(ele, 1)
	fmt.Println("Part A:", part_a)

	key := 811589153
	for i, _ := range ele {
		ele[i] *= key
	}
	part_b := Mix(ele, 10)
	fmt.Println("Part B:", part_b)
}
