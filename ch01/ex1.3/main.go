package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	printArgs()
	fmt.Printf("%vns elapsed\n", time.Since(start).Nanoseconds())
	fmt.Println(strings.Join(os.Args, " "))
}

func printArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func printArgs2() {
	fmt.Println(strings.Join(os.Args, " "))
}
