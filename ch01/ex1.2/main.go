package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		fmt.Printf("%v\t%v\n", k, v)
	}
}
