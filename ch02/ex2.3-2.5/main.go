package main

import (
	"fmt"

	"github.com/tesujiro/ProgrammingLanguageGo/ch02/ex2.3-2.5/popcount"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(popcount.PopCount(uint64(i)))
		fmt.Println(popcount.PopCount2(uint64(i)))
		fmt.Println(popcount.PopCount3(uint64(i)))
	}
}
