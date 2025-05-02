package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service"
	proposalhdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_proposal_service"
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

profile := r.Group("/profile")
{
	profile.POST("/", middleware.AuthMiddleware(), handler.CreateProfile)
	profile.PUT("/", middleware.AuthMiddleware(), handler.UpdateProfile)
	profile.GET("/:user_id", middleware.AuthMiddleware(), handler.GetProfile)
}

review := api.Group("/review")
{
	review.POST("/", middleware.AuthMiddleware(), handler.SubmitReview)
	review.GET("/:freelancer_id", middleware.AuthMiddleware(), handler.GetFreelancerReviews)
}

proposal := r.Group("/proposal")
{
	proposal.POST("/proposal", middleware.AuthMiddleware(), proposalhdlr.CreateProposalHandler)
	proposal.GET("/proposal/:id", middleware.AuthMiddleware(), proposalhdlr.GetProposalByIDHandler)
	proposal.PUT("/proposal/:id", middleware.AuthMiddleware(), proposalhdlr.UpdateProposalHandler)
	proposal.GET("/proposals", middleware.AuthMiddleware(), proposalhdlr.ListProposalsHandler)
	proposal.POST("/template", middleware.AuthMiddleware(), proposalhdlr.SaveTemplateHandler)
	proposal.GET("/templates/:freelancer_id", middleware.AuthMiddleware(), proposalhdlr.GetTemplatesHandler)
	
}

	}

	return r
}
