package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}

	headerSlice := strings.Split(authHeader, " ")
	if len(headerSlice) != 2 || headerSlice[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}
	return headerSlice[1], nil
}
