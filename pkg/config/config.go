package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	LoginPort   int
	LobbyPort   int
	NetworkPort int
	MainPort    int
	RedisURL    string
	ServerIP    string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &Config{
		LoginPort:   getEnvAsInt("LOGIN_PORT", 10881),
		LobbyPort:   getEnvAsInt("LOBBY_PORT", 20202),
		NetworkPort: getEnvAsInt("NETWORK_PORT", 20201),
		MainPort:    getEnvAsInt("MAIN_PORT", 20200),
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
		ServerIP:    getEnv("SERVER_IP", "localhost"),
	}
}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func getEnvAsInt(key string, defaultValue int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
