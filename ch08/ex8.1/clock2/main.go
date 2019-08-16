package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var port int
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.IntVar(&port, "port", 8010, "port number")
	err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("argument parse err:%v\n", err)
		os.Exit(1)
	}

	envTz := os.Getenv("TZ")
	if envTz == "" {
		fmt.Println("env TZ is not set or nil")
		os.Exit(1)
	}
	loc, err := time.LoadLocation(envTz)
	if err != nil {
		fmt.Println("load location error")
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, loc) // handle connections concurrently
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
