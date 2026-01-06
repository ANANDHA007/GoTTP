package main

import (
	"GoTTP/transport"
	"fmt"
	"log"
	"net"
)

func main() {

	tcpListener, err := transport.NewTcpListener(":8000")
	if err != nil {
		log.Fatal(err)
	}
	err1 := tcpListener.Start(func(conn net.Conn) {
		defer conn.Close()
		fmt.Println("Remote Connection Received:", conn.RemoteAddr())
	})

	if err1 != nil {
		log.Fatal(err1)
	}
}
