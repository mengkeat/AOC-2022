package main

import (
	"fmt"
	"os"
	"strings"
)

type Set struct {
	in [53]bool
}

func getIndex(c byte) byte {
	var i byte = 0
	if c >= 'a' && c <= 'z' {
		i = c - 'a'
	} else {
		i = c - 'A' + 26
	}
	return i + 1
}

func (s *Set) Add(a []byte) {
	for _, c := range a {
		if c == 13 {
			continue
		}
		s.in[getIndex(c)] = true
	}
}

func (s Set) exist(c byte) bool {
	if c == 13 {
		return false
	}
	return s.in[getIndex(c)]
}

func main() {
	bytes, _ := os.ReadFile("input03.txt")
	lines := strings.Split(string(bytes), "\n")

	var sum int32 = 0
	for _, ln := range lines {
		s := &Set{}
		l := len(ln)
		s.Add([]byte(ln[:l/2]))
		for _, c := range ln[l/2:] {
			if s.exist(byte(c)) {
				sum += int32(getIndex(byte(c)))
				break
			}
		}
	}
	fmt.Println("Part A:", sum)

	sum = 0
	for i := 0; i < len(lines)/3; i++ {
		s1, s2, s3 := &Set{}, &Set{}, &Set{}
		s1.Add([]byte(lines[i*3]))
		s2.Add([]byte(lines[i*3+1]))
		s3.Add([]byte(lines[i*3+2]))

		for j := 1; j < 53; j++ {
			if s1.in[j] && s2.in[j] && s3.in[j] {
				sum += int32(j)
			}
		}
	}
	fmt.Println("Part B:", sum)
}
