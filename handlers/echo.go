package handlers

import "net/http"

func Echo(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("str")
	w.Write([]byte("Echo text: " + str))
}