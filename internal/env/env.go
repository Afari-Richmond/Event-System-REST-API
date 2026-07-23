package env

import (
	"os"
	"strconv"
)

func GetEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intvalue, err := strconv.Atoi(value); err == nil {
			return intvalue
		}
	}
	return defaultValue
}
