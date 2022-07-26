package main

import (
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func dumpEnv() (envVars map[string]string) {
	envVars = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envVars[pair[0]] = pair[1]
	}
	return
}
