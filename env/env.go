// Package env provides helpers to access environment
// variables. There are utility functions to get env
// vars as Integers, Boolean etc.
package env

import (
	"fmt"
	"os"
	"strconv"
)

// Get returns the value of the specified key
// If the key is not present in environment,
// it returns the specified fallback
func Get(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		value = fallback
	}

	return value
}

// MustGet surely returns the value of the
// specified key. If the key is not found in
// environment, an error is returned.
func MustGet(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s not set in env", key)
	}

	return value, nil
}

// GetInt returns the value as an Integer.
// The fallback is returned if the key is not
// present in environment or is an invalid integer.
func GetInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	intvalue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intvalue
}

// MustGetInt returns the value only if the key
// is set in environment and is an Integer
func MustGetInt(key string) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return 0, fmt.Errorf("%s not set in env", key)
	}

	intvalue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s is not an integer", key)
	}

	return intvalue, nil
}
