package config

import (
	"log"
	"os"
)

var (
	Port      string
	JWTSecret string
)

func InitConfig() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		log.Println("PORT not set in .env, using default 8080")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}
}
