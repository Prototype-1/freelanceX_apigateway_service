package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler"
	"github.com/Prototype-1/freelanceX_apigateway_service/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
			auth.POST("/oauth", handler.OAuth)
			auth.POST("/select-role", middleware.AuthMiddleware(), handler.SelectRole)
			auth.POST("/logout", middleware.AuthMiddleware(), handler.Logout)

		}
	}

	return r
}
