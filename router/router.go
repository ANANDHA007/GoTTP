package router

import "GoTTP/handler"

type Router struct {
	routes map[string]handler.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]handler.HandlerFunc),
	}
}

func (r *Router) Add(method, path string, h handler.HandlerFunc) {
	r.routes[method+" "+path] = h
}

func (r *Router) Match(method, path string) (handler.HandlerFunc, bool) {
	h, ok := r.routes[method+" "+path]
	return h, ok
}
