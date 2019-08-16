package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8010")
	if err != nil {
		log.Fatal("net.Dial: ", err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done: closed by server")
		done <- struct{}{} // signal the main goroutine
	}()
	in := make(chan string)
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			in <- input.Text()
		}
	}()
LOOP:
	for {
		select {
		case text := <-in:
			n, err := fmt.Fprintln(conn, text)
			if n == 0 {
				log.Println("done: closed by client")
				break LOOP
			}
			if err != nil {
				log.Fatal("write error:", err)
				break LOOP
			}
		case <-done:
			break LOOP
		}
	}

	//conn.(*net.TCPConn).CloseWrite()
	//conn.(*net.TCPConn).CloseRead()
	conn.Close()
}
