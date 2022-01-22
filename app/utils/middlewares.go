package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

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

// AuthenticationRequiredMiddleware will check if the user has valid permissions
func AuthenticationRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		// Parse request header and extract the token
		authHeader := request.Header.Get("Authorization")
		tokenValue := strings.Split(authHeader, "Bearer ")
		if len(tokenValue) != 2 {
			ErrorResponseWriter(rw, http.StatusBadRequest, "Error in parsing token", nil)
			return
		}

		tokenString := tokenValue[1]
		
		claims := &UserClaim{}
		token, err := ParseToken(tokenString, claims)
		if err != nil {
			logger.Write("[ERROR]: Error in parsing token - %s", err)
			if err == jwt.ErrSignatureInvalid {
				ErrorResponseWriter(rw, http.StatusUnauthorized, "Invalid signature", nil)
				return
			}

			vErr, _ := err.(*jwt.ValidationError)
			if vErr.Errors == jwt.ValidationErrorExpired {
				ErrorResponseWriter(rw, http.StatusUnauthorized, "Token expired", nil)
				return
			}
			
			ErrorResponseWriter(rw, http.StatusBadRequest, "Error in parsing token", nil)
			return
		}

		if !token.Valid {
			logger.Write("Invalid token")
			ErrorResponseWriter(rw, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		next.ServeHTTP(rw, request)
	})
}
