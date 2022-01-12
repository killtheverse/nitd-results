package utils

import (
	"io/ioutil"

	"github.com/golang-jwt/jwt"
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
func CreateSignedToken(claims *UserClaim) (string, error) {
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
func ParseToken(tokenString string, claims *UserClaim) (*jwt.Token, error) {
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
