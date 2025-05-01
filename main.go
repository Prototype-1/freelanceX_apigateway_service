package main

import (
	"log"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/router"
)

func main() {
	redis.InitRedis()

	client.InitUserServiceClients()

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
