package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(username string) (string, error) {
	secretKey := os.Getenv("SECRET_STRING")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Set token expiration (e.g., 15 minutes)
	})

	return claims.SignedString([]byte(secretKey))
}

func parseToken(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_STRING")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(secretKey), nil
	})
}

func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	type Token struct {
		Token string `json:"token"`
	}
	type User struct {
		username   string `json:"username"`
		expiryDate string `json:"expiryDate"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Token{}
	err := decoder.Decode(&params)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("something went wrong parsing json  %v", err))
	}
	// check token
	tokenString := params.Token
	token, tokenErr := parseToken(tokenString)

	if tokenErr != nil {
		respWithError(w, 400, fmt.Sprintf("something went wrong in parsing the token %v", tokenErr))
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		respWithError(w, 400, fmt.Sprintf("something went wrong in parsing the token"))
		return
	}
	username := claims["username"].(string) // Assuming "username" is the claim key

	popo := User{
		username: username,
	}

	if err != nil {
		respWithError(w, 400, "Something Went wrong parsing the json")
		return
	}

	respWithJSON(w, 200, popo)
}
