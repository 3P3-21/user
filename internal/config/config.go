package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "BUILD_VERSION"

type App struct {
	Addr        string
	Port        int
	Secret      string
	PostgresDSN string
}

type PostgresDB struct {
	DNS string
}

type Config struct {
	App App
}

// New returns a new Config struct
func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		App: App{
			Addr:        getEnv("APP_ADDR", "localhost"),
			Port:        getEnvAsInt("APP_PORT", 3000),
			Secret:      getEnv("APP_AUTH_SECRET", "secret"),
			PostgresDSN: getEnv("POSTGRES_DSN", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"),
		},
	}, nil
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
