package middleware

import "net/http"

func RequireAuth(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if user is authenticated
		token := r.Header.Get("Authorization")
		if token != "Bearer Secret" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}
		f.ServeHTTP(w, r)
	})
}
