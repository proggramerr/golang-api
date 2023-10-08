package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type PostgreSQLConfig struct {
	PostgreHost     string
	PostgrePort     int
	PostgreUser     string
	PostgrePassword string
	PostgreDbName   string
}

type Config struct {
	PostgreSQL PostgreSQLConfig
	DebugMode  bool
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &Config{
		PostgreSQL: PostgreSQLConfig{
			PostgreHost:     getEnv("POSTGRES_HOST", ""),
			PostgrePort:     getEnvAsInt("POSTGRES_PORT", 5432),
			PostgreUser:     getEnv("POSTGRES_USER", ""),
			PostgrePassword: getEnv("POSTGRES_PASSWORD", ""),
			PostgreDbName:   getEnv("POSTGRES_DB", ""),
		},
		DebugMode: getEnvAsBool("debug", true),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
