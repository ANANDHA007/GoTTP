package main

import (
	"GoTTP/connection"
	"GoTTP/http"
	"GoTTP/transport"
	"GoTTP/worker"
	"bufio"
	"fmt"
	"log"
	"net"
)

const SOH = byte(0x01)

func main() {

	tcpListener, err := transport.NewTcpListener(":8000")
	if err != nil {
		log.Fatal(err)
	}

	wp := worker.NewWorkerPool(4) // 4 worker goroutines

	err1 := tcpListener.Start(func(conn net.Conn) {
		defer conn.Close()
		reader := bufio.NewReader(conn)
		for {
			req, err := connection.ReadAndParseRequest(reader)
			if err != nil {
				return
			}

			respCh := make(chan *http.Response, 1)
			wp.JobQueue <- &worker.Job{Req: req, RespCh: respCh}

			resp := <-respCh

			handleRequest(conn, resp)
		}
	})

	if err1 != nil {
		log.Fatal(err1)
	}
}

func handleRequest(conn net.Conn, res *http.Response) {
	fmt.Printf("Connection id %v %d", conn, res.Body)
}
