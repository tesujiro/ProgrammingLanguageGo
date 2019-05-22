package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tesujiro/ProgrammingLanguageGo/ch02/ex2.1-2.2/tempconv"
)

func main() {
	// excercise 2.1
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))     // "212°F"
	fmt.Println(tempconv.CToK(tempconv.BoilingC))     // "212°F"

	// excercise 2.2
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}
