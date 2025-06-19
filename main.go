// @title FreelanceX API Gateway
// @version 1.0
// @description This is the API Gateway for FreelanceX microservices.
// @host freelancex.goxtrace.shop
// @BasePath /

package main

import (
	"log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
	_ "github.com/Prototype-1/freelanceX_apigateway_service/docs"
	"github.com/joho/godotenv"
	"github.com/Prototype-1/freelanceX_apigateway_service/config"
	"github.com/Prototype-1/freelanceX_apigateway_service/kafka"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/router"
	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis"
	projecthdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_project.crm_service"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func startMetricsServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Serving Prometheus metrics on :2112/metrics")
	err := http.ListenAndServe(":2112", mux)
	if err != nil {
		log.Fatalf("Metrics server error: %v", err)
	}
}

func main() {
go startMetricsServer()
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	err = kafka.InitKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatalf("Kafka init failed: %v", err)
	}

	config.InitConfig()
	redis.InitRedis()
	client.InitUserServiceClients()
	client.InitProposalServiceCLient()
	client.InitCrmServiceClients()
	client.InitTimeServiceClients()
	client.InitMessageServiceClient()
	client.InitInvoiceServiceClients()

	clientHandler := &projecthdlr.ClientHandler{ClientClient: client.ClientClient}
	projectHandler := &projecthdlr.ProjectHandler{ProjectClient: client.ProjectClient}	

	r := router.SetupRouter(clientHandler, projectHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
