package main

import (
	"fmt"
	"os"
)

func GetEnvOrDefault(name string, def string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	} else {
		return def
	}
}

func GetEnvOrError(name string) (string, error) {
	if val, ok := os.LookupEnv(name); ok {
		return val, nil
	} else {
		return "", fmt.Errorf("environment variable \"%v\" not set", name)
	}
}