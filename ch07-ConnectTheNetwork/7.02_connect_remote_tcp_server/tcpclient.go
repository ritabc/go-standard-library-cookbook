package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

func createServer(address string) http.Server {
	return http.Server{
		Addr:    address,
		Handler: StringServer("Hello Gophers!\n"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	// Connect with plain TCP
	// Dial produces the Conn, which serves as a reference to the opened socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "GET / HTTP/1.1\nHost: localhost:7070\n\n")

	if err != nil {
		panic(err)
	}

	// Use Scanner to obtain response, which works because of brake lines
	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}
