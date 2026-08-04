package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBDriver  string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBUrl     string
	JWTSecret string
	GIN_MODE  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	jwtSecret := getEnv("JWT_SECRET", "your_super_secret_jwt_key")

	dbDriver := getEnv("DB_DRIVER", "postgres")
	dbHost := getEnv("PGHOST", getEnv("DB_HOST", ""))
	dbUser := getEnv("POSTGRES_USER", getEnv("PGUSER", getEnv("DB_USER", "postgres")))
	dbPass := getEnv("POSTGRES_PASSWORD", getEnv("PGPASSWORD", getEnv("DB_PASSWORD", "")))
	dbName := getEnv("POSTGRES_DB", getEnv("PGDATABASE", getEnv("DB_NAME", "railway")))
	dbPort := getEnv("PGPORT", getEnv("DB_PORT", "5432"))
	
	// Support Railway DATABASE_URL & DATABASE_PUBLIC_URL
	dbUrl := getEnv("DATABASE_URL", getEnv("DATABASE_PUBLIC_URL", getEnv("DB_URL", "")))

	return &Config{
		Port:      getEnv("PORT", "6080"),
		DBDriver:  dbDriver,
		DBHost:    dbHost,
		DBUser:    dbUser,
		DBPass:    dbPass,
		DBName:    dbName,
		DBPort:    dbPort,
		DBUrl:     dbUrl,
		JWTSecret: jwtSecret,
		GIN_MODE:  getEnv("GIN_MODE", "release"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
