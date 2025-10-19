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

// JsonEncode marshals v to JSON bytes.
func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// WriteJSON writes pre-encoded JSON bytes to the ResponseWriter with the
// provided HTTP status code and appropriate Content-Type header.
// It assumes `b` is already a valid JSON encoding. It returns any write error.
func WriteJSON(w http.ResponseWriter, status int, b []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err := w.Write(b)
	return err
}
