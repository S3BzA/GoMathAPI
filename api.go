package main

import (
	"github.com/S3BzA/GoMathAPI/handlers"
	"github.com/S3BzA/GoMathAPI/middleware"
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

	crud := http.NewServeMux()
	crud.HandleFunc("POST /create",handlers.Create)
	crud.HandleFunc("GET /read/{id}",handlers.Read)
	crud.HandleFunc("PUT /update/{id}",handlers.Update)
	crud.HandleFunc("DELETE /delete/{id}",handlers.Delete)

	root_router.Handle("/math/", http.StripPrefix("/math", mathematics))
	root_router.Handle("/crud/", http.StripPrefix("/crud", crud))

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Authentication,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: stack(root_router),
	}
	log.Printf("Server started %s", s.addr)
	return server.ListenAndServe()
}
