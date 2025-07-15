package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GetPathParam extracts a path part by index
func getPathParam(r *http.Request, index int) (string, error) {
	parts := strings.Split(r.URL.Path, "/")
	if index >= len(parts) || parts[index] == "" {
		return "", fmt.Errorf("invalid path: missing value at index %d", index)
	}
	return parts[index], nil
}

// GetPathInt safely converts a path param to int
func GetPathInt(r *http.Request, index int) (int, error) {
	str, err := getPathParam(r, index)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value: %s", str)
	}

	return num, nil
}

// GetPathString returns a string from path
func GetPathString(r *http.Request, index int) (string, error) {
	return getPathParam(r, index)
}
