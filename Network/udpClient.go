package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	fmt.Println("Connecting to server at ", service)
	conn, err := net.Dial("udp", service)

	if err != nil {
		fmt.Println("Could not resolve udp address or connect to it on ", service)
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to server at ", service)

	for i := 0; ; {
		i += 1
		time.Sleep(1000 * time.Millisecond)
		n, err := conn.Write([]byte(string(i)))
		if err != nil {
			fmt.Println("error writing data to server", service)
			fmt.Println(err)
			return
		}

		if n > 0 {
			fmt.Println("Wrote", n, " bytes to server at ", service)
		}
	}

}
