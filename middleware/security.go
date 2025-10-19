package middleware

import "net/http"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api_key := r.Header.Get("api-key")

		if api_key == "" || api_key != "supersecretvalue" {
			err := http.StatusUnauthorized
			http.Error(w, "Invalid API Key", err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
