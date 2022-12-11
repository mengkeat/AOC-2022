package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	items         []int
	op            func(int) int
	test          int
	true_monkey   int
	false_monkey  int
	inspect_count int
}

func monkey_turn(monkeys []Monkey, i int, mod int) {
	curr := &monkeys[i]
	curr.inspect_count += len(curr.items)
	for _, x := range curr.items {
		new_worry := 0
		if mod < 0 {
			new_worry = curr.op(x) / 3
		} else {
			new_worry = curr.op(x) % mod
		}
		dest_monkey := -1
		if new_worry%curr.test == 0 {
			dest_monkey = curr.true_monkey
		} else {
			dest_monkey = curr.false_monkey
		}
		monkeys[dest_monkey].items = append(monkeys[dest_monkey].items, new_worry)
	}
	curr.items = curr.items[:0]
}

func compute_score(monkeys []Monkey, num_rounds int, mod int) int {
	for r := 0; r < num_rounds; r++ {
		for i := 0; i < len(monkeys); i++ {
			monkey_turn(monkeys, i, mod)
		}
	}
	insp := make([]int, 0)
	for i := 0; i < len(monkeys); i++ {
		insp = append(insp, monkeys[i].inspect_count)
	}
	sort.Ints(insp)
	return insp[len(monkeys)-1] * insp[len(monkeys)-2]
}

func main() {
	// var orig_monkeys = []Monkey{
	// 	{[]int{79, 98}, func(x int) int { return x * 19 }, 23, 2, 3, 0},
	// 	{[]int{54, 65, 75, 74}, func(x int) int { return x + 6 }, 19, 2, 0, 0},
	// 	{[]int{79, 60, 97}, func(x int) int { return x * x }, 13, 1, 3, 0},
	// 	{[]int{74}, func(x int) int { return x + 3 }, 17, 0, 1, 0},
	// }

	var orig_monkeys = []Monkey{
		{[]int{65, 58, 93, 57, 66}, func(x int) int { return x * 7 }, 19, 6, 4, 0},
		{[]int{76, 97, 58, 72, 57, 92, 82}, func(x int) int { return x + 4 }, 3, 7, 5, 0},
		{[]int{90, 89, 96}, func(x int) int { return x * 5 }, 13, 5, 1, 0},
		{[]int{72, 63, 72, 99}, func(x int) int { return x * x }, 17, 0, 4, 0},
		{[]int{65}, func(x int) int { return x + 1 }, 2, 6, 2, 0},
		{[]int{97, 71}, func(x int) int { return x + 8 }, 11, 7, 3, 0},
		{[]int{83, 68, 88, 55, 87, 67}, func(x int) int { return x + 2 }, 5, 2, 1, 0},
		{[]int{64, 81, 50, 96, 82, 53, 62, 92}, func(x int) int { return x + 5 }, 7, 3, 0, 0},
	}

	monkeys := make([]Monkey, len(orig_monkeys))
	for i, m := range orig_monkeys {
		monkeys[i] = m
		monkeys[i].items = make([]int, len(m.items))
		copy(monkeys[i].items, m.items)
	}
	fmt.Println("Part A:", compute_score(monkeys, 20, -1))

	// mod := 23 * 19 * 13 * 17
	mod := 19 * 3 * 13 * 17 * 2 * 11 * 5 * 7
	fmt.Println("Part B:", compute_score(orig_monkeys, 10000, mod))
}
