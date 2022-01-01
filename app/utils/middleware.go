package utils

import (
	"net/http"
	"time"
	
	logger "github.com/killtheverse/nitd-results/app/logging"
)

// LoggingMiddleware will log all requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(rw, req)
		logger.Write("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

// JSONContentTypeMiddleware will add the json content type header for all endpoints
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json; charset=UTF-8")
		next.ServeHTTP(rw, r)
	})
}

