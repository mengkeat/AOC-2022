package main

import (
	"fmt"
	"os"
	"strings"
)

func get_start(s []rune, win int) int {
	set := make([]int, 26)
	var i int = 0
	for i < len(s) {
		set[s[i]-'a']++
		b := true
		for j := i - win + 1; j <= i && j >= 0; j++ {
			b = b && (set[s[j]-'a'] == 1)
		}
		if b == true && i >= win-1 {
			break
		}
		if i >= win-1 {
			set[s[i-win+1]-'a']--
		}
		i++
	}
	return i + 1
}

func main() {
	bytes, _ := os.ReadFile("input06.txt")
	ln := strings.Split(string(bytes), "\n")[0]

	fmt.Println("Part A:", get_start([]rune(ln), 4))
	fmt.Println("Part B:", get_start([]rune(ln), 14))
}
