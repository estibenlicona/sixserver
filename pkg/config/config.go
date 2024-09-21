package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sixserver/pkg/types"
	"strconv"
)

func Load() *types.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &types.Config{
		ServerIP:    getEnv("SERVER_IP", "localhost"),
		LoginPort:   getEnvAsInt("LOGIN_PORT", 10881),
		LobbyPort:   getEnvAsInt("LOBBY_PORT", 20202),
		NetworkPort: getEnvAsInt("NETWORK_PORT", 20201),
		MainPort:    getEnvAsInt("MAIN_PORT", 20200),
		CipherKey:   getEnv("CIPHER_KEY", ""),
		Redis: types.RedisConfig{
			Addr:     getEnv("REDIS_URL", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
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
