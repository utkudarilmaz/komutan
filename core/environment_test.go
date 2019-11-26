package core

import (
	"os"
	"testing"
)

func TestGetEnvOrError(t *testing.T) {
	trueKey := "TEST"
	defaultValue := "test-value"

	os.Setenv(trueKey, defaultValue)
	value, err := GetEnvOrError(trueKey)
	if err != nil {
		t.Errorf(
			"%s environment variable setted but error occured when key getted",
			trueKey,
		)
	}
	if value != "test-value" {
		t.Errorf(
			"%s key's value not true. It should be %s",
			trueKey,
			defaultValue,
		)
	}

	falseKey := "FALSE_VALUE"
	_, err = GetEnvOrError(falseKey)
	if err == nil {
		t.Errorf(
			"%s environment variable not setted but function don't return error",
			falseKey,
		)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	defaultValue := "default-value"
	settedValue := "setted-value"
	key := "TEST_DEFAULT"

	value := GetEnvOrDefault(key, defaultValue)
	if value != "default-value" {
		t.Errorf(
			"%s environment variable not setted but default value not returned",
			key,
		)
	}

	os.Setenv(key, settedValue)
	value = GetEnvOrDefault(key, defaultValue)
	if value != "setted-value" {
		t.Errorf(
			"%s environment variable setted but returned value is not true",
			key,
		)
	}
}
