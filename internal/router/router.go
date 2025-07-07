package router

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	projecthdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_project.crm_service"
	proposalhdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_proposal_service"
	timeTrackerHdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_timeTracker_service"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service"
	messagehdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_message.notification_service"
	invoicehdlr "github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_invoice.payment_service"
	"github.com/Prototype-1/freelanceX_apigateway_service/middleware"
	"github.com/Prototype-1/freelanceX_apigateway_service/websocket"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	clientHandler *projecthdlr.ClientHandler,
	projectHandler *projecthdlr.ProjectHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
})
r.GET("/readyz", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ready"})
})

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	 hub := websocket.NewHub()
    go hub.Run()
	messageHandler := messagehdlr.NewMessageHandler(client.MessageClient)

    wsGroup := r.Group("/ws")
    wsGroup.GET("/messages", middleware.AuthMiddleware(), websocket.ServeWS(hub, client.MessageClient))

paymentHandler := invoicehdlr.NewPaymentHandler(client.PaymentClient)

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

		portfolio := r.Group("/portfolio", middleware.AuthMiddleware())
{
	portfolio.POST("/create", handler.CreatePortfolio)
	portfolio.GET("/get/:freelancer_id", handler.GetPortfolio)
	portfolio.DELETE("/delete/:portfolio_id", handler.DeletePortfolio)
}

profile := r.Group("/profile", middleware.AuthMiddleware())
{
	profile.POST("/create", handler.CreateProfile)
	profile.PUT("/update", handler.UpdateProfile)
	profile.GET("/get/:user_id", handler.GetProfile)
}

review := api.Group("/review", middleware.AuthMiddleware())
{
	review.POST("/submit", handler.SubmitReview)
	review.GET("/get/:freelancer_id", handler.GetFreelancerReviews)
}

proposal := r.Group("/proposal", middleware.AuthMiddleware())
{
	proposal.POST("/create", proposalhdlr.CreateProposalHandler)
	proposal.GET("/get/:id", proposalhdlr.GetProposalByIDHandler)
	proposal.PUT("/update/:id", proposalhdlr.UpdateProposalHandler)
	proposal.GET("/listall", proposalhdlr.ListProposalsHandler)
	proposal.POST("/template/save", proposalhdlr.SaveTemplateHandler)
	proposal.GET("/templates/:freelancer_id", proposalhdlr.GetTemplatesHandler)
	
}

client := api.Group("/clients", middleware.AuthMiddleware())
		{
			client.POST("/create", clientHandler.CreateClientHandler)
			client.GET("/get/:id", clientHandler.GetClientHandler)
			client.PUT("/update/:id", clientHandler.UpdateClientHandler)
			client.DELETE("/delete/:id", clientHandler.DeleteClientHandler)
		}

		project := api.Group("/projects", middleware.AuthMiddleware())
		{
			project.POST("/create", projectHandler.CreateProjectHandler)
			project.GET("/get/user/:id", projectHandler.GetProjectsByUserHandler)
			project.GET("/get/project/:id", projectHandler.GetProjectByIdHandler)
			project.GET("/discover/:userId", projectHandler.DiscoverProjectsHandler)
			project.POST("/assign", projectHandler.AssignFreelancerHandler)
			project.PUT("/update/:id", projectHandler.UpdateProjectHandler)
			project.DELETE("/delete/:id", projectHandler.DeleteProjectHandler)
		}

		timeTracker := api.Group("/time-tracker", middleware.AuthMiddleware())
	{
		timeTracker.POST("/logs/create", timeTrackerHdlr.CreateTimeLogHandler)
		timeTracker.GET("/logs/user/:userId", timeTrackerHdlr.GetTimeLogsByUserHandler)
		timeTracker.GET("/logs/project/:project_id", timeTrackerHdlr.GetTimeLogsByProjectHandler)
		timeTracker.PUT("/logs/update/:logId", timeTrackerHdlr.UpdateTimeLogHandler)
		timeTracker.DELETE("/logs/delete/:logId",  timeTrackerHdlr.DeleteTimeLogHandler)
	}

	message := api.Group("/message")
	message.GET("/get/all", middleware.AuthMiddleware(), messageHandler.GetMessages)

	milestone := r.Group("/milestone", middleware.AuthMiddleware())
{
	milestone.POST("/create", invoicehdlr.CreateMilestoneRuleHandler)
	milestone.PUT("/update", invoicehdlr.UpdateMilestoneRuleHandler)
	milestone.GET("/project/:project_id", invoicehdlr.GetMilestonesByProjectIDHandler)
	milestone.GET("/project/:project_id/phase/:phase", invoicehdlr.GetMilestoneByProjectIDAndPhaseHandler)
}

invoiceGroup := r.Group("/invoices", middleware.AuthMiddleware())
{
	invoiceGroup.POST("", invoicehdlr.CreateInvoiceHandler)
	invoiceGroup.GET("/:id", invoicehdlr.GetInvoiceHandler)
	invoiceGroup.GET("/user/:userId", invoicehdlr.GetInvoicesByUserHandler)
	invoiceGroup.GET("/project/:projectId", invoicehdlr.GetInvoicesByProjectHandler)
	invoiceGroup.PUT("/:id/status", invoicehdlr.UpdateInvoiceStatusHandler)
}

payment := r.Group("/payment")
{
	payment.POST("/order", middleware.AuthMiddleware(), paymentHandler.CreatePaymentOrderHandler)
	payment.GET("/checkout", paymentHandler.RazorpayCheckoutPageHandler)
	payment.POST("/verify", paymentHandler.VerifyPaymentHandler)
}

	}
	return r
}