package main

import (
	"log"
	"net/http"
	"github.com/S3BzA/GoMathAPI/handlers"
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
	router := http.NewServeMux()
	router.HandleFunc("GET /echo/{str}", handlers.Echo)

	mathematics := http.NewServeMux()
	mathematics.HandleFunc("GET /add/{nums}", handlers.Add)
	mathematics.HandleFunc("GET /sub/{nums}", handlers.Sub)
	mathematics.HandleFunc("GET /mul/{nums}", handlers.Mul)
	mathematics.HandleFunc("GET /div/{nums}", handlers.Div)
	mathematics.HandleFunc("GET /pow/{nums}", handlers.Pow)
	mathematics.HandleFunc("GET /mod/{nums}", handlers.Mod)

	router.Handle("/math/", http.StripPrefix("/math", mathematics))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
