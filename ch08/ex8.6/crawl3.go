package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/tesujiro/TheGoProgrammingLanguage/ch05/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var maxDepth int

func crawl(url string, depth int) {
	if depth == maxDepth {
		return
	}

	fmt.Println("depth:", depth, "url:", url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	wg := &sync.WaitGroup{}
	for _, url := range list {
		copiedUrl := url
		wg.Add(1)
		go func() {
			crawl(copiedUrl, depth+1)
			wg.Done()
		}()
	}
	wg.Wait()
	return
}

func main() {

	flag.IntVar(&maxDepth, "d", 2, "max crawl depth")
	flag.Parse()

	wg := &sync.WaitGroup{}
	for _, url := range flag.Args() {
		copiedUrl := url
		wg.Add(1)
		go func() {
			crawl(copiedUrl, 0)
			wg.Done()
		}()
	}
	wg.Wait()
}
