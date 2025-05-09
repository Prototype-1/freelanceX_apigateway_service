package handler

import (
	"context"
	"net/http"
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	pb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_timeTracker_service"
	"google.golang.org/grpc/metadata"
)

func withMetadata(c *gin.Context) context.Context {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	md := metadata.Pairs("user_id", userID, "role", role)
	return metadata.NewOutgoingContext(c.Request.Context(), md)
}

func CreateTimeLogHandler(c *gin.Context) {
	var req struct {
		ProjectID string `json:"project_id"`
		TaskName  string `json:"task_name"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Source    string `json:"source"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.CreateTimeLogRequest{
		ProjectId: req.ProjectID,
		TaskName:  req.TaskName,
	StartTime: timestamppb.New(mustParseTime(req.StartTime)),
    EndTime:   timestamppb.New(mustParseTime(req.EndTime)),
		Source:    pb.TimeLogSource(pb.TimeLogSource_value[req.Source]),
	}

	resp, err := client.TimeClient.CreateTimeLog(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetTimeLogsByUserHandler(c *gin.Context) {
	userID := c.Param("user_id")
	projectID := c.Query("project_id")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

var fromTime, toTime *timestamppb.Timestamp
if dateFrom != "" {
    parsedFromTime, err := time.Parse(time.RFC3339, dateFrom)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for 'date_from'"})
        return
    }
    fromTime = timestamppb.New(parsedFromTime)
}

if dateTo != "" {
    parsedToTime, err := time.Parse(time.RFC3339, dateTo)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for 'date_to'"})
        return
    }
    toTime = timestamppb.New(parsedToTime)
}

	grpcReq := &pb.GetTimeLogsByUserRequest{
		UserId:   userID,
		ProjectId: projectID,
		DateFrom: fromTime,
		DateTo:   toTime,
	}

	resp, err := client.TimeClient.GetTimeLogsByUser(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetTimeLogsByProjectHandler(c *gin.Context) {
	projectID := c.Param("project_id")

	grpcReq := &pb.GetTimeLogsByProjectRequest{
		ProjectId: projectID,
	}

	resp, err := client.TimeClient.GetTimeLogsByProject(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateTimeLogHandler(c *gin.Context) {
	var req struct {
		LogID    string `json:"log_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.UpdateTimeLogRequest{
		LogId:    req.LogID,
	StartTime: timestamppb.New(mustParseTime(req.StartTime)),
EndTime:   timestamppb.New(mustParseTime(req.EndTime)),
	}

	resp, err := client.TimeClient.UpdateTimeLog(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteTimeLogHandler(c *gin.Context) {
	logID := c.Param("log_id")

	grpcReq := &pb.DeleteTimeLogRequest{
		LogId: logID,
	}

	resp, err := client.TimeClient.DeleteTimeLog(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func mustParseTime(timeStr string) time.Time {
    parsedTime, err := time.Parse(time.RFC3339, timeStr)
    if err != nil {
        panic(err) 
    }
    return parsedTime
}