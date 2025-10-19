package utils

import (
	"encoding/json"
	"net/http"
)

type MathResult struct {
	Result float64 `json:"result"`
}

type MathError struct {
	Error string `json:"error"`
}

func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func WriteJSON(w http.ResponseWriter, status int, b []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err := w.Write(b)
	return err
}
