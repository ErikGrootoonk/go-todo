package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath string
	Port   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "todos.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{DBPath: dbPath, Port: port}
}
