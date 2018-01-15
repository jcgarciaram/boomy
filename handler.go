package main

import (
	"errors"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jcgarciaram/boomy/utils"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

		secret := utils.GetJWTSecret()

		if len(secret) == 0 {
			return nil, errors.New("Auth0 Client Secret Not Set")
		}

		return []byte(secret), nil
	},

	Extractor: extractToken,
	Debug:     false,
})

func extractToken(r *http.Request) (string, error) {

	var token string

	// get Authorization Header
	auth := r.Header.Get("Authorization")

	token = strings.Split(auth, " ")[1]

	return token, nil

}
