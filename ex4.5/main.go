package main

import "fmt"

func main() {
	uniq := func(sa []string) []string {
		var prev string
		index := 0
		for i := 0; i < len(sa); i++ {
			current := (sa)[i]
			if i == 0 || current != prev {
				sa[index] = current
				index++
				prev = current
			}
		}
		return sa[:index]
	}

	list := []string{"hello", "world", "world", "hello"}
	//slice := list[:]
	fmt.Println(list)
	fmt.Println(uniq(list))
	//fmt.Println(slice)
}
