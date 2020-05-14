package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go" // go get -u github.com/dgrijalva/jwt-go
)

var mySigningKey = []byte("mysupersecretphrase")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	
	claims := token.Claims.(jwt.MapClaims)
	
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	tokenString, err := generateJWT()

	if err != nil {
		fmt.Println("error generating token string")
		return
	}
	fmt.Println(tokenString)
}