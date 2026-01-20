package main

import (
	"GoTTP/transport"
	"fmt"
	"io"
	"log"
	"net"
)

const SOH = byte(0x01)

func main() {

	tcpListener, err := transport.NewTcpListener(":8000")
	if err != nil {
		log.Fatal(err)
	}
	err1 := tcpListener.Start(func(conn net.Conn) {
		defer conn.Close()
		fmt.Println("Accepted from", conn.RemoteAddr())

		buf := make([]byte, 64)
		readPos := 0
		writePos := 0

		for {

			if writePos == len(buf) {
				if readPos > 0 {
					copy(buf, buf[readPos:writePos])
					writePos -= readPos
					readPos = 0
				} else {
					newBuf := make([]byte, len(buf)*2)
					copy(newBuf, buf)
					buf = newBuf
				}
			}
			n, err := conn.Read(buf[writePos:])

			if err != nil {
				if err == io.EOF {
					fmt.Println("client closed connection:", conn.RemoteAddr())
					return
				}
				log.Println("read error:", err)
				return
			}

			writePos += n
			for {
				msgEnd := FindMessageEnd(buf[readPos:writePos])
				fmt.Println("Message Size :", msgEnd)
				fmt.Println("Buffer Appended String :", string(buf))
				fmt.Println("Value of write ", writePos)
				if msgEnd < 0 {
					break
				}
				msg := buf[readPos : readPos+msgEnd]
				fmt.Printf("Received message from %s: %q\n", conn.RemoteAddr(), string(msg))
				readPos += msgEnd
			}

			if readPos > 0 {
				copy(buf, buf[readPos:writePos])
				writePos -= readPos
				readPos = 0
			}
		}
	})

	if err1 != nil {
		log.Fatal(err1)
	}
}

func FindMessageEnd(buf []byte) int {
	for i := 0; i < len(buf)-3; i++ {
		if buf[i] == '\r' && buf[i+1] == '\n' && buf[i+2] == '\r' && buf[i+3] == '\n' {
			return i + 4
		}
	}
	return -1
}
