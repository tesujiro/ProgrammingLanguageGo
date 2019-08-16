package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8010")
	if err != nil {
		log.Fatal("Server Listen:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Server Accept:", err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	in := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			in <- input.Text()
		}
		// NOTE: ignoring potential errors from input.Err()
	}()
	wg := &sync.WaitGroup{}
LOOP:
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Server:Timeout")
			break LOOP
		case text := <-in:
			wg.Add(1)
			go func() {
				echo(c, text, 1*time.Second)
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
