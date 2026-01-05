package transport

import (
	"fmt"
	"net"
)

type TcpListner struct {
	Addr string
}

func NewTcpListener(tcp string) (*TcpListner, error) {

	tcplistener := TcpListner{
		Addr: tcp,
	}
	return &tcplistener, nil
}

func (tcp *TcpListner) Start() (bool, error) {

	tcpln, err := net.Listen("tcp", tcp.Addr)

	if err != nil {
		return false, err
	}

	for {
		conn, err := tcpln.Accept()
		defer conn.Close()
		if err != nil {
			return false, err
		}
		fmt.Printf("Received from: %v ", conn.RemoteAddr())
	}

}
