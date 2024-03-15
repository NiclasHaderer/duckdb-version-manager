package utils

import "os"

func EnvIsTruthy(env string) bool {
	value := os.Getenv(env)
	return value == "true" || value == "1"
}
