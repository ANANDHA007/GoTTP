package handler

import "GoTTP/context"

type HandlerFunc func(ctx *context.Context)

type Handler interface {
	ServeHTTP(ctx *context.Context)
}

func (f HandlerFunc) ServeHTTP(ctx *context.Context) {
	f(ctx)
}
