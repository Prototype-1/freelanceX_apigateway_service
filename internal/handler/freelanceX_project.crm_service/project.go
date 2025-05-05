package handlers

import (
	"context"
	"net/http"

	projectpb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_project.crm_service/project"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectClient projectpb.ProjectServiceClient
}

func NewProjectHandler(projectClient projectpb.ProjectServiceClient) *ProjectHandler {
	return &ProjectHandler{ProjectClient: projectClient}
}

func (h *ProjectHandler) CreateProjectHandler(c *gin.Context) {
	var req projectpb.CreateProjectRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProjectClient.CreateProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *ProjectHandler) GetProjectsByUserHandler(c *gin.Context) {
	userId := c.Param("userId")

	res, err := h.ProjectClient.GetProjectsByUser(context.Background(), &projectpb.GetProjectsByUserRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Projects)
}

func (h *ProjectHandler) GetProjectByIdHandler(c *gin.Context) {
	projectId := c.Param("id")

	res, err := h.ProjectClient.GetProjectById(context.Background(), &projectpb.GetProjectByIdRequest{ProjectId: projectId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) DiscoverProjectsHandler(c *gin.Context) {
	userId := c.Param("userId")

	res, err := h.ProjectClient.DiscoverProjects(context.Background(), &projectpb.DiscoverProjectsRequest{UserId: userId})
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

	res, err := h.ProjectClient.AssignFreelancer(context.Background(), &req)
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

	res, err := h.ProjectClient.UpdateProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProjectHandler) DeleteProjectHandler(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ProjectClient.DeleteProject(context.Background(), &projectpb.DeleteProjectRequest{ProjectId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
