package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
func createSignedToken(claims *UserClaim) (string, error) {
	privateKey, err := ioutil.ReadFile("key.rsa")
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err 
	}
	return tokenString, nil
}

// parseToken will parse a token from a string
func parseToken(tokenString string, claims *UserClaim) (*jwt.Token, error) {
	publicKey, err := ioutil.ReadFile("key.rsa.pub")
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
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
		tokenString, err := createSignedToken(claims)
		if err != nil {
			logger.Write("[ERROR]: Error in creating token - %s", err)
			utils.ResponseWriter(rw, http.StatusInternalServerError, "Error occured", nil)
			return
		} else {
			logger.Write("Credentials validated")
			response := map[string]string {
				"token": tokenString,
			}
			utils.ResponseWriter(rw, http.StatusOK, "Signed in", response)
		}
		
	} else {
		logger.Write("Invalid credentials")
		utils.ResponseWriter(rw, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}
}


func Refresh(rw http.ResponseWriter, request *http.Request) {
	// Parse request header and extract the token
	authHeader := request.Header.Get("Authorization")
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	
	claims := &UserClaim{}
	token, err := parseToken(tokenString, claims)
	if err != nil {
		logger.Write("[ERROR]: Error in parsing token - %s", err)
		if err == jwt.ErrSignatureInvalid {
			utils.ResponseWriter(rw, http.StatusUnauthorized, "Invalid signature", nil)
			return
		}

		vErr, _ := err.(*jwt.ValidationError)
		if vErr.Errors == jwt.ValidationErrorExpired {
			utils.ResponseWriter(rw, http.StatusUnauthorized, "Token expired", nil)
			return
		}
		
		utils.ResponseWriter(rw, http.StatusBadRequest, "", nil)
		return
	}

	if !token.Valid {
		logger.Write("Invalid token")
		utils.ResponseWriter(rw, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		logger.Write("Token has not expired")
		utils.ResponseWriter(rw, http.StatusBadRequest, "Token has not expired", nil)
		return
	}

	// Create a new token with renewed expiration time
	expirationTime := time.Now().Add(1*time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	tokenString, err = createSignedToken(claims)
	if err != nil {
		logger.Write("[ERROR]: Error in creating token - %s", err)
		utils.ResponseWriter(rw, http.StatusInternalServerError, "Error occured", nil)
		return
	} else {
		logger.Write("Credentials validated")
		response := map[string]string {
			"token": tokenString,
		}
		utils.ResponseWriter(rw, http.StatusOK, "Token refeshed", response)
	}
	logger.Write("claims: %v", claims)
}