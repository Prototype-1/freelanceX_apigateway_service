package client

import (
	"log"
	"os"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authPb "github.com/Prototype-1/freelanceX_user_service/proto/auth"
//	profilePb "github.com/Prototype-1/freelanceX_user_service/proto/profile"
	//portfolioPb "github.com/Prototype-1/freelanceX_user_service/proto/portfolio"
//	reviewPb "github.com/Prototype-1/freelanceX_user_service/proto/review"
)

var (
	AuthClient      authPb.AuthServiceClient
//	ProfileClient   profilePb.ProfileServiceClient
//	PortfolioClient portfolioPb.PortfolioServiceClient
//	ReviewClient    reviewPb.ReviewServiceClient
)

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
	//ProfileClient = profilePb.NewProfileServiceClient(conn)
	//PortfolioClient = portfolioPb.NewPortfolioServiceClient(conn)
	//ReviewClient = reviewPb.NewReviewServiceClient(conn)

	log.Println("Connected to User Service via gRPC at", userGrpcAddr)
}
