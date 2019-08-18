package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//response := make(chan io.ReadCloser, len(os.Args))
	response := make(chan string)

	wg := sync.WaitGroup{}
	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			//fmt.Println(url)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatalf("http new request error: %v", err)
				return
			}
			req.Cancel = ctx.Done()
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("http client do error: %v", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				resp.Body.Close()
				log.Fatalf("http status code is not OK: %v", resp.StatusCode)
				return
			}
			defer resp.Body.Close()
			fmt.Println("finished;", url)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			response <- string(bodyBytes)

		}(arg)
	}

	body := <-response

	cancel() // cancel other requests
	wg.Wait()
	close(response)
	fmt.Println("body:", body)
}
