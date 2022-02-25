package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/utils"
)

// swagger:route POST /auth/signin/ authentication signIn
// Sign in a user
// Consumes:
// - application/json
// 
// Produces:
// - application/json
// 
// Schemes: http, https
//
// Responses:
//	default: ErrorResponse
//	200: ErrorResponse
func SignIn(rw http.ResponseWriter, request *http.Request) {
	// Parse request body to get the credentials
	var creds utils.Credentials

	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		utils.ErrorResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: %v", err)
		return
	}

	// Get the expected password
	expectedPassword := os.Getenv("ADMIN_PASSWORD")

	// If the username and password are correct, create and return a signed token
	// If not, then return a response stating the same
	if creds.Username == "admin" && expectedPassword == creds.Password {
		// Expiration time for the token
		expirationTime := time.Now().Add(time.Minute*15)
		
		// Create claim which includes username and expiration time
		claims := &utils.UserClaim{
			&jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
			utils.UserInfo{Username: creds.Username},
		}

		// Generate token
		tokenString, err := utils.CreateSignedToken(claims)
		if err != nil {
			logger.Write("[ERROR]: Error in creating token - %s", err)
			utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured", nil)
			return
		} else {
			logger.Write("Credentials validated")
			response := map[string]string {
				"token": tokenString,
			}
			utils.ErrorResponseWriter(rw, http.StatusOK, "Signed in", response)
		}
		
	} else {
		logger.Write("Invalid credentials")
		utils.ErrorResponseWriter(rw, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}
}
