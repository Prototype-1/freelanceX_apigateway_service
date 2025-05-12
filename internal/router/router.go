package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service"
	proposalhdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_proposal_service"
	projecthdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_project.crm_service"
	timeTrackerHdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_timeTracker_service"
	"github.com/Prototype-1/freelanceX_apigateway_service/middleware"
	"github.com/Prototype-1/freelanceX_apigateway_service/websocket"
)

func SetupRouter(
	clientHandler *projecthdlr.ClientHandler,
	projectHandler *projecthdlr.ProjectHandler,
) *gin.Engine {
	r := gin.Default()

	 hub := websocket.NewHub()
    go hub.Run()

    wsGroup := r.Group("/ws")
    wsGroup.GET("/messages", middleware.AuthMiddleware(), websocket.ServeWS(hub))

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
			auth.POST("/oauth", handler.OAuth)
			auth.POST("/select-role",  handler.SelectRole)
			auth.GET("/me", middleware.AuthMiddleware(), handler.GetMe)
			auth.POST("/logout", middleware.AuthMiddleware(), handler.Logout)

		}

		portfolio := r.Group("/portfolio")
{
	portfolio.POST("/create", middleware.AuthMiddleware(), handler.CreatePortfolio)
	portfolio.GET("/get/:freelancer_id", middleware.AuthMiddleware(), handler.GetPortfolio)
	portfolio.DELETE("/delete/:portfolio_id", middleware.AuthMiddleware(), handler.DeletePortfolio)
}

profile := r.Group("/profile")
{
	profile.POST("/create", middleware.AuthMiddleware(), handler.CreateProfile)
	profile.PUT("/update", middleware.AuthMiddleware(), handler.UpdateProfile)
	profile.GET("/get/:user_id", middleware.AuthMiddleware(), handler.GetProfile)
}

review := api.Group("/review")
{
	review.POST("/submit", middleware.AuthMiddleware(), handler.SubmitReview)
	review.GET("/get/:freelancer_id", middleware.AuthMiddleware(), handler.GetFreelancerReviews)
}

proposal := r.Group("/proposal")
{
	proposal.POST("/create", middleware.AuthMiddleware(), proposalhdlr.CreateProposalHandler)
	proposal.GET("/get/:id", middleware.AuthMiddleware(), proposalhdlr.GetProposalByIDHandler)
	proposal.PUT("/update/:id", middleware.AuthMiddleware(), proposalhdlr.UpdateProposalHandler)
	proposal.GET("/listall", middleware.AuthMiddleware(), proposalhdlr.ListProposalsHandler)
	proposal.POST("/template/save", middleware.AuthMiddleware(), proposalhdlr.SaveTemplateHandler)
	proposal.GET("/templates/:freelancer_id", middleware.AuthMiddleware(), proposalhdlr.GetTemplatesHandler)
	
}

client := api.Group("/clients")
		{
			client.POST("/create", middleware.AuthMiddleware(), clientHandler.CreateClientHandler)
			client.GET("/get/:id", middleware.AuthMiddleware(), clientHandler.GetClientHandler)
			client.PUT("/update/:id", middleware.AuthMiddleware(), clientHandler.UpdateClientHandler)
			client.DELETE("/delete/:id", middleware.AuthMiddleware(), clientHandler.DeleteClientHandler)
		}

		project := api.Group("/projects")
		{
			project.POST("/create", middleware.AuthMiddleware(), projectHandler.CreateProjectHandler)
			project.GET("/get/user/:id", middleware.AuthMiddleware(), projectHandler.GetProjectsByUserHandler)
			project.GET("/get/project/:id", middleware.AuthMiddleware(), projectHandler.GetProjectByIdHandler)
			project.GET("/discover/:userId", middleware.AuthMiddleware(), projectHandler.DiscoverProjectsHandler)
			project.POST("/assign", middleware.AuthMiddleware(), projectHandler.AssignFreelancerHandler)
			project.PUT("/update/:id", middleware.AuthMiddleware(), projectHandler.UpdateProjectHandler)
			project.DELETE("/delete/:id", middleware.AuthMiddleware(), projectHandler.DeleteProjectHandler)
		}

		timeTracker := api.Group("/time-tracker")
	{
		timeTracker.POST("/logs/create", middleware.AuthMiddleware(), timeTrackerHdlr.CreateTimeLogHandler)
		timeTracker.GET("/logs/user/:userId", middleware.AuthMiddleware(), timeTrackerHdlr.GetTimeLogsByUserHandler)
		timeTracker.GET("/logs/project/:projectId", middleware.AuthMiddleware(), timeTrackerHdlr.GetTimeLogsByProjectHandler)
		timeTracker.PUT("/logs/update/:logId", middleware.AuthMiddleware(), timeTrackerHdlr.UpdateTimeLogHandler)
		timeTracker.DELETE("/logs/delete/:logId", middleware.AuthMiddleware(), timeTrackerHdlr.DeleteTimeLogHandler)
	}

	}
	return r
}