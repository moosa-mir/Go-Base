package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetPathParam extracts a path part by index
func getPathParam(r *http.Request, index int) (string, error) {
	parts := strings.Split(r.URL.Path, "/")
	if index >= len(parts) || parts[index] == "" {
		return "", fmt.Errorf("invalid path: missing value at index %d", index)
	}
	return parts[index], nil
}

func GetQueryParam(r *http.Request, key string) string {
	query := r.URL.Query()
	value := query.Get(key)
	return value
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

// GenerateTimeIntervalFromEpoch calculates the time interval (in seconds) between now and the Unix epoch.
func GenerateTimeIntervalFromEpoch() float32 {
	// Define the Unix epoch (1970-01-01 00:00:00 UTC)
	epoch := time.Unix(0, 0)

	// Calculate the duration between now and the Unix epoch
	duration := time.Since(epoch) // Returns a time.Duration

	// Convert the duration to seconds and return as a float32
	return float32(duration.Seconds())
}

func GetHashPassword(password string) string {
	// Use a fixed salt (not recommended for security-sensitive applications)
	fixedSalt := "my-fixed-salt"

	// Combine the password with the fixed salt
	input := password + fixedSalt

	// Generate a SHA-256 hash
	hash := sha256.Sum256([]byte(input))

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash[:])
}
