package server

import (
	"GoTTP/connection"
	"GoTTP/context"
	"GoTTP/handler"
	"GoTTP/router"
	"GoTTP/transport"
	"GoTTP/worker"
	"bufio"
	"net"
)

type HttpServer struct {
	router *router.Router
	pool   *worker.WorkerPool
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: router.NewRouter(),
		pool:   worker.NewWorkerPool(5),
	}
}

func (s *HttpServer) GET(path string, h handler.HandlerFunc)    { s.router.Add("GET", path, h) }
func (s *HttpServer) POST(path string, h handler.HandlerFunc)   { s.router.Add("POST", path, h) }
func (s *HttpServer) PUT(path string, h handler.HandlerFunc)    { s.router.Add("PUT", path, h) }
func (s *HttpServer) DELETE(path string, h handler.HandlerFunc) { s.router.Add("DELETE", path, h) }

func (s *HttpServer) ListenAndServe(addr string) error {
	listener, _ := transport.NewTcpListener(addr)
	return listener.Start(func(conn net.Conn) {
		defer conn.Close()
		reader := bufio.NewReader(conn)
		for {
			req, err := connection.ReadAndParseRequest(reader)
			if err != nil {
				return
			}

			h, ok := s.router.Match(req.Method, req.Path)
			if !ok {
				return
			}
			ctx := context.New(req, conn)
			s.pool.Submit(h, ctx)
			ctx.WriteResponse()
		}
	})
}
