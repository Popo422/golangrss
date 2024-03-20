package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization header is passed")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed Header no apiKey header placed")
	}
	return vals[1], nil
}
