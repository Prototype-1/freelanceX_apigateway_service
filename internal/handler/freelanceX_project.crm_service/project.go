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
    
    switch v := role.(type) {
    case string:
        return v
    case *string:
        if v != nil {
            return *v
        }
        return ""
    default:
        return ""
    }
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
	userId := c.Param("id")

	if !validateUUID(userId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId format"})
		return
	}

	role := getRoleFromContext(c)
	if role != "client" && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: only clients or admins can access this endpoint"})
		return
	}

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

		if !validateUUID(projectId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid projectId format"})
		return
	}

	role := getRoleFromContext(c)
	if role != "client" && role != "admin" && role != "freelancer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))

	res, err := h.ProjectClient.GetProjectById(ctx, &projectpb.GetProjectByIdRequest{ProjectId: projectId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) DiscoverProjectsHandler(c *gin.Context) {
    fmt.Println("DEBUG: All context keys:")
    for key, value := range c.Keys {
        fmt.Printf("  %s: %v (type: %T)\n", key, value, value)
    }
    
    fmt.Println("DEBUG: Request headers:")
    for key, values := range c.Request.Header {
        fmt.Printf("  %s: %v\n", key, values)
    }
    
    req := projectpb.DiscoverProjectsRequest{
        UserId: c.Param("userId"),
    }
    
    fmt.Printf("DEBUG: UserId from param: '%s'\n", req.UserId)
    
    role := getRoleFromContext(c)
    fmt.Printf("DEBUG: Role extracted: '%s'\n", role)
    
    if role == "" {
        c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Only freelancers are allowed to discover projects",
    "code": "role_forbidden",
		})
        return
    }
    
    if role != "freelancer" {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": fmt.Sprintf("unauthorized: expected 'freelancer', got '%s'", role),
        })
        return
    }
    
    if !validateUUID(req.UserId) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId format"})
        return
    }
    
    fmt.Printf("DEBUG: Creating context with role: '%s'\n", role)
    ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", role))
    
    fmt.Println("DEBUG: Calling backend service...")
    res, err := h.ProjectClient.DiscoverProjects(ctx, &req)
    if err != nil {
        fmt.Printf("DEBUG: Backend service error: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    fmt.Println("DEBUG: Backend service call successful")
    c.JSON(http.StatusOK, res) 
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

	if !validateUUID(req.ProjectId) || !validateUUID(req.FreelancerId) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format for projectId or freelancerId"})
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
