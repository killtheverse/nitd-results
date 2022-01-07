package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/utils"
)

type UserInfo struct {
	Username string
}

type UserClaim struct {
	*jwt.StandardClaims
	UserInfo
}

type Credentials struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

// createSignedToken will create and return a signed token based on the username
func createSignedToken(claim UserClaim) (string, error) {
	privateKey, err := ioutil.ReadFile("key.rsa")
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodRS256)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err 
	}
	return tokenString, nil
}



func SignIn(rw http.ResponseWriter, request *http.Request) {
	// Parse request body to get the credentials
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		utils.ResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: %v", err)
		return
	}

	// Get the expected password
	expectedPassword := os.Getenv("ADMIN_PASSWORD")

	// If the username and password are correct, create and return a signed token
	// If not, then return a response stating the same
	if creds.Username == "admin" && expectedPassword == creds.Password {
		// Expiration time for the token
		expirationTime := time.Now().Add(time.Minute*1)
		
		// Create claim which includes username and expiration time
		claims := &UserClaim{
			&jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
			UserInfo{creds.Username},
		}

		// Generate token
		tokenString, err := createSignedToken(*claims)
		if err != nil {
			logger.Write("[ERROR]: %s", err)
			utils.ResponseWriter(rw, http.StatusInternalServerError, "Error occured", nil)
			return
		} else {
			logger.Write("Credentials validated")

			//Set the client's cookie for "token" as the same as JWT created
			http.SetCookie(rw, &http.Cookie{
				Name: "token",
				Value: tokenString,
				Expires: expirationTime,
				HttpOnly: true,
			})
			utils.ResponseWriter(rw, http.StatusOK, "Signed in", nil)
		}
		
	} else {
		logger.Write("Invalid credentials")
		utils.ResponseWriter(rw, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}
}
