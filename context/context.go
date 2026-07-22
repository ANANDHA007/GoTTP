package context

import (
	"GoTTP/http"
	"fmt"
	"net"
)

type Context struct {
	Request *http.Request
	Conn    net.Conn
	Params  map[string]string
	Status  int
	headers map[string]string
	body    []byte
}

func New(req *http.Request, conn net.Conn) *Context {
	return &Context{
		Request: req,
		Conn:    conn,
		Params:  make(map[string]string),
		headers: make(map[string]string),
		body:    []byte{},
		Status:  200,
	}
}

func (c *Context) String(status int, body string) {
	c.Status = status
	c.headers["Content-Type"] = "text/plain; charset=utf-8"
	c.body = []byte(body)
}

func (c *Context) WriteResponse() error {
	if c.Status == 0 {
		c.Status = 200
	}

	c.headers["Content-Length"] = fmt.Sprintf("%d", len(c.body))

	fmt.Fprintf(c.Conn, "HTTP/1.1 %d OK\r\n", c.Status)

	for key, value := range c.headers {
		fmt.Fprintf(c.Conn, "%s: %s\r\n", key, value)
	}

	fmt.Fprint(c.Conn, "\r\n")

	_, err := c.Conn.Write(c.body)
	return err
}
