package handlers

import (
	"context"
	"net/http"
	projectpb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_project.crm_service/project"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"fmt"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	ProjectClient projectpb.ProjectServiceClient
}

func NewProjectHandler(projectClient projectpb.ProjectServiceClient) *ProjectHandler {
	return &ProjectHandler{ProjectClient: projectClient}
}

func getRoleFromContext(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		return "" 
	}
	return role.(string)
}

func validateUUID(clientId string) bool {
	_, err := uuid.Parse(clientId)
	return err == nil
}

func (h *ProjectHandler) CreateProjectHandler(c *gin.Context) {
	var req projectpb.CreateProjectRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := getRoleFromContext(c)
	fmt.Println("Role from header:", role)
	if role != "client" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: only clients can create projects"})
		return
	}

	if !validateUUID(req.ClientId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clientId format"})
		return
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))
	res, err := h.ProjectClient.CreateProject(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *ProjectHandler) GetProjectsByUserHandler(c *gin.Context) {
	userId := c.Param("userId")

	role := getRoleFromContext(c)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))

	res, err := h.ProjectClient.GetProjectsByUser(ctx, &projectpb.GetProjectsByUserRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Projects)
}

func (h *ProjectHandler) GetProjectByIdHandler(c *gin.Context) {
	projectId := c.Param("id")

	role := getRoleFromContext(c)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))

	res, err := h.ProjectClient.GetProjectById(ctx, &projectpb.GetProjectByIdRequest{ProjectId: projectId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) DiscoverProjectsHandler(c *gin.Context) {
	userId := c.Param("userId")

	role := getRoleFromContext(c)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))

	res, err := h.ProjectClient.DiscoverProjects(ctx, &projectpb.DiscoverProjectsRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Projects)
}

func (h *ProjectHandler) AssignFreelancerHandler(c *gin.Context) {
	var req projectpb.AssignFreelancerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role := getRoleFromContext(c)
	if role != "client" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: only clients can assign freelancers"})
		return
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))
	res, err := h.ProjectClient.AssignFreelancer(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) UpdateProjectHandler(c *gin.Context) {
	id := c.Param("id")
	var req projectpb.UpdateProjectRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ProjectId = id

	role := getRoleFromContext(c)
	if role != "client" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: only clients can update projects"})
		return
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))
	res, err := h.ProjectClient.UpdateProject(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) DeleteProjectHandler(c *gin.Context) {
	id := c.Param("id")

	role := getRoleFromContext(c)
	if role != "admin" && role != "client" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: only admins or clients can delete projects"})
		return
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))
	res, err := h.ProjectClient.DeleteProject(ctx, &projectpb.DeleteProjectRequest{ProjectId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
