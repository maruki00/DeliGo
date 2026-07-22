package handlers

import (
	"errors"
	"net/http"

	"github.com/maruki00/deligo/internal/analytic/requests"
	"github.com/maruki00/deligo/internal/analytic/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FeedbackHandler struct {
	svc service.FeedbackService
}

func NewFeedbackHandler(svc service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{svc: svc}
}

func (h *FeedbackHandler) CreateFeedback(c *gin.Context) {
	var req requests.CreateFeedbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.SubmitFeedback(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record customer feedback application tracking."})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *FeedbackHandler) GetProductAnalytics(c *gin.Context) {
	productID := c.Param("product_id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product UUID verification key parameter must be explicitly passed."})
		return
	}

	res, err := h.svc.GetAnalyticsByProductID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Analytics metric aggregates record historical trace target missing."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
