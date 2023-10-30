package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey  extracts an api key
// the headers of an http request
// example
// Authorization: ApiKey {insert api key here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return " ", errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed auth header")
	}
	return vals[1], nil
}
