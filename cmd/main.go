package main

import (
	"GoTTP/context"
	"GoTTP/server"
)

func main() {

	http := server.NewHttpServer()
	http.GET("/hello", HelloHandler)
	http.ListenAndServe(":8080")

}

func HelloHandler(ctx *context.Context) {
	ctx.String(200, "Hello World")
}
