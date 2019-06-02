package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bs []byte) {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}

func reverseUTF8(bs []byte) {
	for i := 0; i < len(bs); {
		_, size := utf8.DecodeRune(bs[i:])
		if size > 1 {
			reverse(bs[i : i+size])
		}
		i += size
	}
	reverse(bs)
}

func main() {
	list := []byte("Hello, World! こんにちは世界")
	fmt.Printf("%s\n", list)
	reverseUTF8(list)
	fmt.Printf("%s\n", list)
}
