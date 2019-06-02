package main

import "fmt"

func main() {
	uniq := func(sa *[]string) {
		var temp []string
		var prev string
		for i := 0; i < len(*sa); i++ {
			current := (*sa)[i]
			if i == 0 || current != prev {
				temp = append(temp, current)
				prev = current
			}
		}
		*sa = temp
	}

	list := []string{"hello", "world", "world", "hello"}
	fmt.Println(list)
	uniq(&list)
	fmt.Println(list)
}
