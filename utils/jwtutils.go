package utils

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"time"
)

var (
	secret = "script penates onion potence spinning exocrine"
)

type MyCustomClaims struct {
	CustomStruct interface{} `json:"custom_struct"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return secret
}

//GenerateJWT generates JWT web token
func GenerateJWT(customStruct interface{}) (string, error) {

	mySigningKey := []byte(secret)

	// Create the Claims
	claims := MyCustomClaims{
		CustomStruct: customStruct,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    "demo-park",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)

}

// GetCustomStruct looks at the token passed in the request and returns the custom struct stored in request
func GetCustomStruct(r *http.Request) (interface{}, error) {

	// Get user context from request
	user := r.Context().Value("user")

	// Get the map of claims from user context
	mapClaims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error reading user map of claims from JWT")
	}

	return mapClaims["custom_struct"], nil

}
