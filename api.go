package main

import (
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
	router := http.NewServeMux()
	router.HandleFunc("/echo/{str}", Echo)
	
	mathematics := http.NewServeMux()
	mathematics.HandleFunc("/add", Add)
	mathematics.HandleFunc("/sub", Add)
	mathematics.HandleFunc("/mul", Add)
	mathematics.HandleFunc("/div", Add)
	mathematics.HandleFunc("/pow", Add)
	mathematics.HandleFunc("/mod", Add)

	router.Handle("/math/", http.StripPrefix("/math", mathematics))
	
	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}