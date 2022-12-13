package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	isLeaf bool
	val    int
	elems  []Node
}

func newLeaf(v int) Node       { return Node{true, v, []Node{}} }
func singletonNode(v int) Node { return Node{false, 0, []Node{newLeaf(v)}} }

func newList(s string, i int) (Node, int) {
	temp := 0
	node := Node{}
	for i < len(s) {
		if s[i] == '[' {
			n, nexti := newList(s, i+1)
			node.elems = append(node.elems, n)
			i = nexti
		} else if s[i] == ']' {
			if s[i-1] != '[' {
				node.elems = append(node.elems, newLeaf(temp))
			}
			return node, i + 1
		} else if s[i] == ',' {
			node.elems = append(node.elems, newLeaf(temp))
			temp = 0
		} else {
			temp += (temp * 10) + int(s[i]-'0')
		}
		i++
	}
	return node, i + 1
}

// Returns n2-n1
func cmp(n1, n2 Node) int {
	i := 0
	for i < len(n1.elems) && i < len(n2.elems) {
		c := 0
		if n1.elems[i].isLeaf && n2.elems[i].isLeaf {
			c = n2.elems[i].val - n1.elems[i].val
		} else if !n1.elems[i].isLeaf && !n2.elems[i].isLeaf {
			c = cmp(n1.elems[i], n2.elems[i])
		} else {
			if n1.elems[i].isLeaf {
				c = cmp(singletonNode(n1.elems[i].val), n2.elems[i])
			}
			if n2.elems[i].isLeaf {
				c = cmp(n1.elems[i], singletonNode(n2.elems[i].val))
			}
		}
		if c != 0 {
			return c
		}
		i++
	}
	return len(n2.elems) - len(n1.elems)
}

type Signal []Node

func (s Signal) Len() int           { return len(s) }
func (s Signal) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Signal) Less(i, j int) bool { return cmp(s[i], s[j]) >= 0 }

func main() {
	bytes, _ := os.ReadFile("input13.txt")
	lines := strings.Split(string(bytes), "\n")

	d1, _ := newList("[[2]]", 0)
	d2, _ := newList("[[6]]", 0)
	all_nodes := Signal{d1, d2}

	part_a := 0
	for i := 0; i < len(lines); i = i + 3 {
		n1, _ := newList(lines[i], 0)
		n2, _ := newList(lines[i+1], 0)
		all_nodes = append(all_nodes, n1, n2)
		c := cmp(n1, n2)
		if c >= 0 {
			part_a += i/3 + 1
		}
	}
	fmt.Println("Part A:", part_a)

	part_b := 1
	sort.Sort(all_nodes)
	for i, s := range all_nodes {
		if cmp(s, d1) == 0 || cmp(s, d2) == 0 {
			part_b *= i + 1
		}
	}
	fmt.Println("Part B:", part_b)
}
