package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func fatalOnPanic() {
	if p := recover(); p != nil {
		log.Fatal(p)
	}
}

func handleConn(conn net.Conn) {
	if _, err := io.Copy(conn, conn); err != nil {
		log.Println("failed to copy: ", err)
	}

	conn.Close()
}

func main() {
	defer fatalOnPanic()

	var port string
	flag.StringVar(&port, "port", "9090", "Port to listen")

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("listen: %v", err))
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(fmt.Sprintf("accept: %v", err))
		}

		go handleConn(conn)
	}
}
