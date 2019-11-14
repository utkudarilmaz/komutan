package core

import (
	"errors"
	"os"
)

// GetEnvOrError is getting environment variable from system. If
// environment variable's length is diffirent from 0 then return environment
// variable value and nil, otherwise empty string and an error.
func GetEnvOrError(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", errors.New(key + " environment variable not setted.")
	}
	return value, nil
}

// GetEnvOrDefault getting the environment variable value if that is setted,
// otherwise return gived default value.
func GetEnvOrDefault(key, variable string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return variable
	}
	return value
}
