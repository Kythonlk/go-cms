package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs details of every request and response time.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request: Method and URL
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Proceed with the next handler
		next.ServeHTTP(w, r)

		// Log the completion time and status
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
