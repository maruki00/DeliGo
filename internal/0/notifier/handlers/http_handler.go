package handlers

import (
	"net/http"
	"strconv"

	"github.com/maruki00/deligo/internal/notifier/models"
	"github.com/maruki00/deligo/internal/notifier/service"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc service.Service
}

func NewHTTPHandler(svc service.Service) *HTTPHandler {
	return &HTTPHandler{svc: svc}
}

func (h *HTTPHandler) GetNotifications(c *gin.Context) {
	userID := c.GetString("user_id")
	notifications, err := h.svc.GetUnreadNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	c.JSON(http.StatusOK, notifications)
}

func (h *HTTPHandler) MarkAsRead(c *gin.Context) {
	userID := c.GetString("user_id")
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	if err := h.svc.MarkNotificationAsRead(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

func (h *HTTPHandler) GetChatHistory(c *gin.Context) {
	orderID := c.Param("order_id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID required"})
		return
	}

	history, err := h.svc.GetChatHistory(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chat history"})
		return
	}
	c.JSON(http.StatusOK, history)
}

func (h *HTTPHandler) UploadFile(c *gin.Context) {
	var fileReq struct {
		ID       string `json:"id" binding:"required"`
		OwnerID  string `json:"owner_id" binding:"required"`
		FileURL  string `json:"file_url" binding:"required"`
		FileType string `json:"file_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&fileReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileModel := models.File{
		ID:       fileReq.ID,
		OwnerID:  fileReq.OwnerID,
		FileURL:  fileReq.FileURL,
		FileType: fileReq.FileType,
	}

	if err := h.svc.UploadFileMetadata(&fileModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process file tracking details"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File registration complete", "data": fileModel})
}
