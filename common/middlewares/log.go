package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Println("Started: ", r.Method, " ", r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		log.Println("Completed ", r.Method, " ", r.URL.Path, " Time taken: ", duration)
	})
}
