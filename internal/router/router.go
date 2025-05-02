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
			auth.GET("/me", middleware.AuthMiddleware(), handler.GetMe)
			auth.POST("/logout", middleware.AuthMiddleware(), handler.Logout)

		}

		portfolio := r.Group("/portfolio")
{
	portfolio.POST("/", middleware.AuthMiddleware(), handler.CreatePortfolio)
	portfolio.GET("/:freelancer_id", middleware.AuthMiddleware(), handler.GetPortfolio)
	portfolio.DELETE("/:portfolio_id", middleware.AuthMiddleware(), handler.DeletePortfolio)
}

	}

	return r
}
