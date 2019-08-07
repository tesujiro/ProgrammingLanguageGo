package main

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/tesujiro/ProgrammingLanguageGo/ex4.12/api"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print(`	xkcd get [number]
	xkcd index [filename]
	xkcd search [filename] [term]
`)
}

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
		comics, err := index()
		if err != nil {
			fmt.Printf("get description error: %v\n", err)
			os.Exit(1)
		}
		err = save(comics, filename)
		if err != nil {
			fmt.Printf("save index error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Indexing completed: %v\n", filename)
	case cmd == "search" && len(args) >= 2:
		filename := args[0]
		//words := args[1:]
		comics, err := load(filename)
		if err != nil {
			fmt.Printf("load index error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loading completed: %v\n", comics)

		/*
			err, results := search(comics, words)
			if err != nil {
				fmt.Printf("search word error: %v\n", err)
				os.Exit(1)
			}
		*/

	default:
		usage()
		os.Exit(1)
	}
}

func index() ([]api.Comic, error) {
	max, err := api.GetMaxNumber()
	if err != nil {
		return nil, err
	}
	fmt.Println("max =", max)
	max = 10 // temporary

	// Get Descriptions
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

	for i := 1; i <= max; i++ {
		comic := <-res
		fmt.Printf("comics[%v]=%v\n", comic.Num-1, comic.Title)
		comics[comic.Num-1] = comic
	}

	return comics, nil
}

/*
func search(comics []api.Comic, words []string) ([]api.Comic, error) {
	found := make(map[int]int)

	for _, comic := range comics {
	}
	return nil,nilk
}
*/

func save(comics []api.Comic, filename string) error {
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
	return nil
}

func load(filename string) ([]api.Comic, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	dec := gob.NewDecoder(fd)
	var comics []api.Comic
	err = dec.Decode(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}
