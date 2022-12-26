package main

import (
	"fmt"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

var StoN = map[byte]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
var NtoS = map[int]byte{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}

func add_digit(d1, d2 byte, ca int) (string, int) {
	n1, n2 := StoN[d1], StoN[d2]
	c := 0
	d_str := string("")
	switch t := n1 + n2 + ca; t {
	case 3, 4, 5:
		c = 1
		d_str = string(NtoS[t-5])
	case -3, -4, -5:
		c = -1
		d_str = string(NtoS[t+5])
	default:
		c = 0
		d_str = string(NtoS[t])
	}
	return d_str, c
}

func add(s1, s2 string) string {
	res := string("")
	carry := 0
	i := 0

	for i = 0; i < Min(len(s1), len(s2)); i++ {
		d_str, c := add_digit(s1[i], s2[i], carry)
		res = res + d_str
		carry = c
	}
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	for i < len(s1) {
		d_str, c := add_digit(s1[i], '0', carry)
		res = res + d_str
		carry = c
		i++
	}
	if carry != 0 {
		res = res + string(NtoS[carry])
	}
	return res
}
func main() {
	lines := LinesFromFile("input25.txt")

	part_a := "0"
	nums := make([]string, len(lines))
	for i, ln := range lines {
		nums[i] = reverse(ln)
		part_a = add(part_a, nums[i])
	}
	fmt.Println("Part A:", reverse(part_a))

}
