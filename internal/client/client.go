package client

import (
	"log"
	"os"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/auth"
	profilePb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/profile"
	portfolioPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/portfolio"
	reviewPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/review"
	proposalPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_proposal_service"
	projectPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_project.crm_service/project"
	clientsPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_project.crm_service/client"
	timePb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_timeTracker_service"
	messagePb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_message.notification_service"
)

var (
	AuthClient      authPb.AuthServiceClient
	ProfileClient   profilePb.ProfileServiceClient
	PortfolioClient portfolioPb.PortfolioServiceClient
	ReviewClient    reviewPb.ReviewServiceClient

	ProposalClient proposalPb.ProposalServiceClient

	ProjectClient projectPb.ProjectServiceClient
	ClientClient  clientsPb.ClientServiceClient

	TimeClient timePb.TimeLogServiceClient

	 MessageClient messagePb.MessageServiceClient
)

	// --- USER SERVICE ---

func InitUserServiceClients() {
	userGrpcAddr := os.Getenv("USER_SERVICE_GRPC_ADDR")
	if userGrpcAddr == "" {
		userGrpcAddr = "localhost:50051"
	}

	conn, err := grpc.NewClient(userGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	AuthClient = authPb.NewAuthServiceClient(conn)
	ProfileClient = profilePb.NewProfileServiceClient(conn)
	PortfolioClient = portfolioPb.NewPortfolioServiceClient(conn)
	ReviewClient = reviewPb.NewReviewServiceClient(conn)

	log.Println("Connected to User Service via gRPC at", userGrpcAddr)
}

	// --- PROPOSAL SERVICE ---

func InitProposalServiceCLient() {
		proposalGrpcAddr := os.Getenv("PROPOSAL_SERVICE_GRPC_ADDR")
		if proposalGrpcAddr == "" {
			proposalGrpcAddr = "localhost:50052"
		}
	
		proposalConn, err := grpc.NewClient(proposalGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to proposal service: %v", err)
		}
	
		ProposalClient = proposalPb.NewProposalServiceClient(proposalConn)
		log.Println("Connected to Proposal Service via gRPC at", proposalGrpcAddr)
}

// --- PROJECT.CRM SERVICE ---

func InitCrmServiceClients() {
	crmGrpcAddr := os.Getenv("CRM_SERVICE_GRPC_ADDR")
	if crmGrpcAddr == "" {
		crmGrpcAddr = "localhost:50053"
	}

	clientConn, err := grpc.NewClient(crmGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create CRM service gRPC client: %v", err)
	}

	ProjectClient = projectPb.NewProjectServiceClient(clientConn)
	ClientClient = clientsPb.NewClientServiceClient(clientConn)

	log.Println("Connected to CRM Service via gRPC at", crmGrpcAddr)
}

// --- TIME LOG SERVICE ---

func InitTimeServiceClients() {
	timeGrpcAddr := os.Getenv("TIMELOG_SERVICE_GRPC_ADDR")
	if timeGrpcAddr == "" {
		timeGrpcAddr = "localhost:50054"
	}

	timeConn, err := grpc.NewClient(timeGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create CRM service gRPC client: %v", err)
	}

	TimeClient = timePb.NewTimeLogServiceClient(timeConn)

	log.Println("Connected to CRM Service via gRPC at", timeGrpcAddr)
}

//---MESSAGE SERVICE---
func InitMessageServiceClient() {
    messageGrpcAddr := os.Getenv("MESSAGE_SERVICE_GRPC_ADDR")
    if messageGrpcAddr == "" {
        messageGrpcAddr = "localhost:50055" 
    }

    conn, err := grpc.NewClient(messageGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to message service: %v", err)
    }

    MessageClient = messagePb.NewMessageServiceClient(conn)

    log.Println("Connected to Message Service via gRPC at", messageGrpcAddr)
}