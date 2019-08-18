package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/tesujiro/TheGoProgrammingLanguage/ch08/ex8.10/links"
)

func crawl(ctx context.Context, url string) []string {
	fmt.Println(url)
	list, err := links.Extract(ctx, url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	wg := sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-c:
			fmt.Println("SIGNAL RECEIVED")
			cancel()
		case <-ctx.Done():
		}
	}()

	// Add command-line arguments to worklist.
	wg.Add(1)
	go func() {
		defer wg.Done()
		worklist <- os.Args[1:]
	}()

	// Create 5 crawler goroutines to fetch each unseen link.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case link, ok := <-unseenLinks:
					if !ok {
						return
					}
					foundLinks := crawl(ctx, link)
					wg.Add(1)
					go func() {
						defer wg.Done()
						select {
						case worklist <- foundLinks:
						case <-ctx.Done():
							fmt.Println("CANNOT write to worklist")
						}
					}()
				case <-ctx.Done():
					fmt.Println("crawler QUIT.")
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
LOOP:
	for {
		select {
		case list, ok := <-worklist:
			if !ok {
				return
			}
			for _, link := range list {
				if !seen[link] {
					select {
					case unseenLinks <- link:
						seen[link] = true
					case <-ctx.Done():
						fmt.Println("The main goroutine CANCELED 1.")
						break LOOP
					}
				}
			}
		case <-ctx.Done():
			fmt.Println("The main goroutine CANCELED 2.")
			break LOOP
		}
	}
	close(unseenLinks)
	wg.Wait()
	close(worklist)
	fmt.Println("The main goroutine finish.")
}
