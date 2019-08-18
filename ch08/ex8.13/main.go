// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

//type client chan<- string // an outgoing message channel
type client struct {
	name    string
	message chan<- string // an outgoing message channel
}

const timeout = 10 * time.Second

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster(ctx context.Context) {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.message <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			var names []string
			for cli := range clients {
				names = append(names, cli.name)
			}
			for cli := range clients {
				cli.message <- fmt.Sprintf("client list: %v", names)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.message)
		case <-ctx.Done():
			//fmt.Println("broadcaster cancel")
			for cli := range clients {
				delete(clients, cli)
				close(cli.message)
			}
			return
		}
	}
}

func handleConn(ctx context.Context, conn net.Conn) {
	wg := sync.WaitGroup{}
	ch := make(chan string) // outgoing client messages

	wg.Add(1)
	go func() {
		defer wg.Done()
		clientWriter(ctx, conn, ch)
		//fmt.Println("handleConn clientWriter goroutine cancel.")
	}()

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	cli := client{name: who, message: ch}
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			conn.Close()
		case <-ctx.Done():
			conn.Close()
		}
		//fmt.Println("handleConn timer goroutine cancel.")
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	// NOTE: ignoring potential errors from input.Err()

	select {
	case leaving <- cli:
		messages <- who + " has left"
	case <-ctx.Done():
	}
	conn.Close()
	wg.Wait()
	fmt.Println("handleConn main goroutine cancel.")
}

func clientWriter(ctx context.Context, conn net.Conn, ch <-chan string) {
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
		case <-ctx.Done():
			//fmt.Println("clientWriter cancel")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	listener, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}

	closeCh := make(chan struct{}, 1) // buffer=1

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-c:
			fmt.Println("SIGNAL RECEIVED")
			cancel()
			closeCh <- struct{}{} // no wait because buffer=1
			listener.Close()
		case <-ctx.Done():
		}
		//fmt.Println("main goroutine cancel")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		broadcaster(ctx)
	}()

LOOP:
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-closeCh:
				break LOOP
			default:
			}
			log.Print(err)
			continue
		}
		wg.Add(1)
		go func() {
			handleConn(ctx, conn)
			wg.Done()
		}()
	}
	wg.Wait()
}
