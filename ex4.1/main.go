package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countBitDiff(c1, c2 [sha256.Size]byte) int {
	if len(c1) != len(c2) {
		panic("different length!")
	}
	count := 0
	for i, _ := range c1 {
		xor := c1[i] ^ c2[i] // bit difference
		count += int(pc[xor])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("diff=%v\n", countBitDiff(c1, c2))

}
