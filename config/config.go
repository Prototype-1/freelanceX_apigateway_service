package config

import (
	"log"
	"os"
)

var (
	Port string
)

func InitConfig() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080" 
		log.Println("PORT not set in .env, using default 8080")
	}
}
