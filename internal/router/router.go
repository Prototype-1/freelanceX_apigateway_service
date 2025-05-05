package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service"
	proposalhdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_proposal_service"
	projecthdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_project.crm_service"
	"github.com/Prototype-1/freelanceX_apigateway_service/middleware"
)

func SetupRouter(
	clientHandler *projecthdlr.ClientHandler,
	projectHandler *projecthdlr.ProjectHandler,
) *gin.Engine {
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
	profile.PUT("/update", middleware.AuthMiddleware(), handler.UpdateProfile)
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

client := api.Group("/clients")
		{
			client.POST("/", middleware.AuthMiddleware(), clientHandler.CreateClientHandler)
			client.GET("/:id", middleware.AuthMiddleware(), clientHandler.GetClientHandler)
			client.PUT("/:id", middleware.AuthMiddleware(), clientHandler.UpdateClientHandler)
			client.DELETE("/:id", middleware.AuthMiddleware(), clientHandler.DeleteClientHandler)
		}

		// Project routes
		project := api.Group("/projects")
		{
			project.POST("/", middleware.AuthMiddleware(), projectHandler.CreateProjectHandler)
			project.GET("/user/:userId", middleware.AuthMiddleware(), projectHandler.GetProjectsByUserHandler)
			project.GET("/:id", middleware.AuthMiddleware(), projectHandler.GetProjectByIdHandler)
			project.GET("/discover/:userId", middleware.AuthMiddleware(), projectHandler.DiscoverProjectsHandler)
			project.POST("/assign", middleware.AuthMiddleware(), projectHandler.AssignFreelancerHandler)
			project.PUT("/:id", middleware.AuthMiddleware(), projectHandler.UpdateProjectHandler)
			project.DELETE("/:id", middleware.AuthMiddleware(), projectHandler.DeleteProjectHandler)
		}

	}
	return r
}
