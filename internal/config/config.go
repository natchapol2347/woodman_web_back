package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct { // TODO: private type
	Host     string `env:"POSTGRESQL_URI"` // TODO: rename to URI
	Name     string `env:"POSTGRESQL_NAME"`
	Username string `env:"POSTGRESQL_USERNAME"`
	Password string `env:"POSTGRESQL_PASSWORD"`
	Port     string `env:"POSTGRESQL_PORT"`
}

func LoadConfig() Database {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := Database{
		Host:     os.Getenv("POSTGRESQL_URI"),
		Name:     os.Getenv("POSTGRESQL_NAME"),
		Username: os.Getenv("POSTGRESQL_USERNAME"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Port:     os.Getenv("POSTGRESQL_PORT"),
	}

	return db
}
