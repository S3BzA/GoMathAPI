package main

import (
	"github.com/S3BzA/GoMathAPI/handlers"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	root_router := http.NewServeMux()
	root_router.HandleFunc("GET /echo/{str}", handlers.Echo)

	mathematics := http.NewServeMux()
	mathematics.HandleFunc("GET /sum/{nums}", handlers.Sum)
	mathematics.HandleFunc("GET /mul/{nums}", handlers.Mul)
	mathematics.HandleFunc("GET /div/{nums}", handlers.Div)
	mathematics.HandleFunc("GET /pow/{nums}", handlers.Pow)
	mathematics.HandleFunc("GET /mod/{nums}", handlers.Mod)

	root_router.Handle("/math/", http.StripPrefix("/math", mathematics))

	server := http.Server{
		Addr:    s.addr,
		Handler: root_router,
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
