package webogo

import (
	"log"
	"net/http"
	"strings"
)

type F map[string]interface{}

type Handler func(*Context)

type Handlers map[string]Handler

type Routers struct {
	router map[string]Handlers
}

func Default() *Routers {
	Routers := Routers{
		router: make(map[string]Handlers),
	}
	return &Routers
}

func (r *Routers) Handle(httpMethod, relativePath string, handler Handler) {
	handlers, ok := r.router[httpMethod]
	if !ok {
		s := make(Handlers)
		r.router[httpMethod] = s
		handlers = s
	}
	_, ok = handlers[relativePath]
	if ok {
		panic("same route")
	}
	handlers[relativePath] = handler
}

func (r *Routers) GET(relativePath string, handler Handler) {
	r.Handle("GET", relativePath, handler)
}

func (r *Routers) POST(relativePath string, handler Handler) {
	r.Handle("POST", relativePath, handler)
}

func (r *Routers) PUT(relativePath string, handler Handler) {
	r.Handle("PUT", relativePath, handler)
}

func (r *Routers) DELETE(relativePath string, handler Handler) {
	r.Handle("DELETE", relativePath, handler)
}

func (r *Routers) Run(port string) {
	if port == "" {
		port = ":8080"
	}
	http.Handle("/",r)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (r *Routers) ServeHTTP(res http.ResponseWriter,req *http.Request) {
	httpMethod := req.Method
	uri := req.RequestURI
	params := strings.Split(uri, "?")
	if len(params) == 0 {
		return
	}
	handlers, ok1 := r.router[httpMethod]
	if !ok1 {
		log.Println("error", req.RemoteAddr)
		return
	}
	handler, ok2 := handlers[params[0]]
	if !ok2 {
		handler404(req, res)
		return
	}
	c := InitContext(req, res)
	handler(&c)
}

func handler404(req *http.Request, res http.ResponseWriter) {
	_, _ = res.Write([]byte("404 Not Found"))
}
