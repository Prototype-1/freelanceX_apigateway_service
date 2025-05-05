package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/Prototype-1/freelanceX_apigateway_service/config"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/router"
	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis"
	projecthdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_project.crm_service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	config.InitConfig()
	redis.InitRedis()
	client.InitUserServiceClients()
	client.InitProposalServiceCLient()
	client.InitCrmServiceClients()

	clientHandler := &projecthdlr.ClientHandler{ClientClient: client.ClientClient}
	projectHandler := &projecthdlr.ProjectHandler{ProjectClient: client.ProjectClient}	

	r := router.SetupRouter(clientHandler, projectHandler)

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
