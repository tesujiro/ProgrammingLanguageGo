package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/tesujiro/ProgrammingLanguageGo/ex4.12/api"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print(`	xkcd get [number]
	xkcd index [filename]
	xkcd search [filename] [term]
`)
}

type Index map[string][]int

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]
	switch {
	case cmd == "get" && len(args) == 1:
	case cmd == "index" && len(args) == 1:
		filename := args[0]
		comics, idx, err := index()
		if err != nil {
			fmt.Printf("get description error: %v\n", err)
			os.Exit(1)
		}
		err = save(comics, idx, filename)
		if err != nil {
			fmt.Printf("save index error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Indexing completed: %v\n", filename)
	case cmd == "search" && len(args) >= 2:
		filename := args[0]
		words := args[1:]
		comics, idx, err := load(filename)
		if err != nil {
			fmt.Printf("load index error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loading completed: %v comics\n", len(comics))

		/*
			// DUMP INDEZX
			for word, idx := range idx {
				fmt.Printf("%v: %v\n", word, idx)
			}
		*/

		results, err := idx.search(words)
		if err != nil {
			fmt.Printf("search word error: %v\n", err)
			os.Exit(1)
		}
		for _, num := range results {
			comic := comics[num]
			fmt.Printf("[%v]:\n%v\n\n", comic.Title, comic.Transcript)
		}

	default:
		usage()
		os.Exit(1)
	}
}

func index() ([]api.Comic, Index, error) {
	// Get max comic number
	max, err := api.GetMaxNumber()
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("max =", max)
	max = 10 // temporary

	// Start Worker for fetching cosmic descriptions
	workers := 5
	comics := make([]api.Comic, max)
	req := make(chan int, workers)
	res := make(chan api.Comic, workers)

	for i := 1; i <= workers; i++ {
		go func() {
			for num := range req {
				comic, err := api.GetDescription(num)
				if err != nil {
					fmt.Printf("Can't get comic %d: %s", num, err)
					os.Exit(1)
				}
				res <- comic
			}
		}()
	}

	for i := 1; i <= max; i++ {
		req <- i
	}
	close(req)

	// Get Descriptions
	for i := 1; i <= max; i++ {
		comic := <-res
		fmt.Printf("comics[%v]=%v\n", comic.Num-1, comic.Title)
		comics[comic.Num-1] = comic
	}

	// Make index from cosmic transcrit
	idx := make(Index)
	for i := 0; i < max; i++ {
		comic := comics[i]
		wordFlag := make(map[string]bool)
		rep := regexp.MustCompile(`[^ 0-9A-Za-z]`)
		text := rep.ReplaceAllString(comic.Transcript, "")
		scanner := bufio.NewScanner(strings.NewReader(text))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := scanner.Text()
			if !wordFlag[word] {
				if _, ok := idx[word]; !ok {
					idx[word] = make([]int, 0)
				}
				idx[word] = append(idx[word], comic.Num)
				wordFlag[word] = true
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}

	return comics, idx, nil
}

func (idx Index) search(words []string) ([]int, error) {
	for _, word := range words {
		if nums, ok := idx[word]; ok {
			return nums, nil
		}
	}
	return nil, nil
}

func save(comics []api.Comic, idx Index, filename string) error {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	enc := gob.NewEncoder(fd)
	err = enc.Encode(comics)
	if err != nil {
		return err
	}
	err = enc.Encode(idx)
	if err != nil {
		return err
	}
	return nil
}

func load(filename string) ([]api.Comic, Index, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	dec := gob.NewDecoder(fd)
	var comics []api.Comic
	err = dec.Decode(&comics)
	if err != nil {
		return nil, nil, err
	}
	var idx Index
	err = dec.Decode(&idx)
	if err != nil {
		return nil, nil, err
	}
	return comics, idx, nil
}
