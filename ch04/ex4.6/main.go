package main

import (
	"fmt"
	"unicode"
)

func squashSpaces(bs []byte) []byte {
	rs := []rune(string(bs))
	index := 0
	adjacent := false
	for _, r := range rs {
		if !unicode.IsSpace(rune(r)) {
			rs[index] = r
			index++
			adjacent = false
		} else if !adjacent {
			rs[index] = ' '
			index++
			adjacent = true
		}
	}
	return []byte(string(rs[:index]))
}

func main() {
	b := []byte("Hello,    World!\nこんにちは、　　世界！\n")
	fmt.Printf("%s", b)
	fmt.Printf("%s", squashSpaces(b))
}
