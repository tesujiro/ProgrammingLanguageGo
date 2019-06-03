// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	var cCount int                  // count Unicode category C
	var lCount int                  // count Unicode category L
	var mCount int                  // count Unicode category M
	var nCount int                  // count Unicode category N
	var pCount int                  // count Unicode category P
	var zCount int                  // count Unicode category Z
	var sCount int                  // count Unicode category S
	categoryList := []struct {
		cat string
		cnt *int
		fnc func(rune) bool
	}{
		{cat: "C", cnt: &cCount, fnc: unicode.IsControl},
		{cat: "L", cnt: &lCount, fnc: unicode.IsLetter},
		{cat: "M", cnt: &mCount, fnc: unicode.IsMark},
		{cat: "N", cnt: &nCount, fnc: unicode.IsNumber},
		{cat: "P", cnt: &pCount, fnc: unicode.IsPunct},
		{cat: "Z", cnt: &zCount, fnc: unicode.IsSpace},
		{cat: "S", cnt: &sCount, fnc: unicode.IsSymbol},
	}
	category := func(r rune) string {
		var cat string
		for _, e := range categoryList {
			if e.fnc(r) {
				cat += e.cat
			}
		}
		return cat
	}

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		for _, e := range categoryList {
			if e.fnc(r) {
				*e.cnt++
			}
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\tcategory\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\t%s\n", c, n, category(c))
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nunicode category\tcount\n")
	for _, e := range categoryList {
		if *e.cnt > 0 {
			fmt.Printf("Category %s\t%d\n", e.cat, *e.cnt)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
