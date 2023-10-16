package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type MongoConfig struct {
	MongoDomain   string
	MongoUser     string
	MongoPassword string
}

type Config struct {
	MongoDB   MongoConfig
	DebugMode bool
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &Config{
		MongoDB: MongoConfig{
			MongoDomain:   getEnv("MONGO_DOMAIN", ""),
			MongoUser:     getEnv("MONGO_USER", ""),
			MongoPassword: getEnv("MONGO_PASSWORD", ""),
		},
		DebugMode: getEnvAsBool("DEBUG", true),
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
