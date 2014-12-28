package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Printf("%s\n", buf)
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
