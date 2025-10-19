package main

import "net/http"
import "github.com/S3BzA/GoMathAPI/mathops"

func Echo(w http.ResponseWriter, r *http.Request) {
		str := r.PathValue("str")
		w.Write([]byte("Echo text: "+str))
}

func Add(w http.ResponseWriter, r *http.Request) {
	
}