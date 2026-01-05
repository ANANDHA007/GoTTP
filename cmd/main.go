package main

import (
	"GoTTP/transport"
	"log"
)

func main() {

	tcpListener, err := transport.NewTcpListener(":8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpListener.Start()
}
