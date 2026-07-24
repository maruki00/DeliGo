package handlers

import (
	"net/http"

	"github.com/maruki00/deligo/internal/courier/domain"
	"github.com/maruki00/deligo/internal/courier/requests"

	"github.com/gin-gonic/gin"
)

type CourierHandler struct {
	svc domain.CourierService
}

func NewCourierHandler(svc domain.CourierService) *CourierHandler {
	return &CourierHandler{svc: svc}
}

func (h *CourierHandler) RegisterCourier(c *gin.Context) {
	var req requests.CourierSetupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.RegisterCourier(req.ID, req.VehicleType, req.Latitude, req.Longitude, req.IsActive); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "courier registered successfully"})
}

func (h *CourierHandler) AcceptOrder(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.AcceptOrder(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "order accepted by courier"})
}

func (h *CourierHandler) ArriveAtRestaurant(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.ArriveAtRestaurant(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "courier arrived at restaurant"})
}

func (h *CourierHandler) StartCourier(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.StartCourier(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delivery transit started"})
}

func (h *CourierHandler) CompleteCourier(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CompleteCourier(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delivery completed successfully"})
}

func (h *CourierHandler) PingLocation(c *gin.Context) {
	var req requests.LocationPingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.TrackLocation(req.CourierID, req.OrderID, req.Latitude, req.Longitude); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "telemetry updated"})
}
