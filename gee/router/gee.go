package router

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	//路由
	router *router
}

// constuctor
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRouter(method string, path string, handler HandlerFunc) {
	engine.router.addRoute(method, path, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRouter("POST", path, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// implement http.Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	engine.router.handle(context)
}
