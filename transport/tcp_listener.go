package transport

import (
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

func (tcp *TcpListner) Start(handler func(net.Conn)) error {

	tcpln, err := net.Listen("tcp", tcp.Addr)

	if err != nil {
		return err
	}

	defer tcpln.Close()

	for {

		conn, err := tcpln.Accept()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}
