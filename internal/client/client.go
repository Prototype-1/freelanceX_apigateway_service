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
)

var (
	AuthClient      authPb.AuthServiceClient
	ProfileClient   profilePb.ProfileServiceClient
	PortfolioClient portfolioPb.PortfolioServiceClient
	ReviewClient    reviewPb.ReviewServiceClient
	ProposalClient proposalPb.ProposalServiceClient
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

	// --- PROPOSAL SERVICE ---
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

