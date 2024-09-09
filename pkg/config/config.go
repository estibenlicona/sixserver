package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	CheckingTCPPort int
	LoginTCPPort    int
}

func Load() Config {
	config := Config{
		CheckingTCPPort: getEnvironmentVariable("CHECKING_PORT"),
		LoginTCPPort:    getEnvironmentVariable("LOGIN_PORT"),
	}
	return config
}

func getEnvironmentVariable(key string) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		panic(fmt.Sprintf("environment variable %s not found", key))
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("invalid value for environment variable %s: %v", key, err))
	}

	return value
}
