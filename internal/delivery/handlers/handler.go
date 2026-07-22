package handlers

import (
	"net/http"

	"github.com/maruki00/deligo/delivery/requests"
	"github.com/maruki00/deligo/delivery/service"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	svc service.DeliveryService
}

func NewDeliveryHandler(svc service.DeliveryService) *DeliveryHandler {
	return &DeliveryHandler{svc: svc}
}

func (h *DeliveryHandler) RegisterCourier(c *gin.Context) {
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

func (h *DeliveryHandler) AcceptOrder(c *gin.Context) {
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

func (h *DeliveryHandler) ArriveAtRestaurant(c *gin.Context) {
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

func (h *DeliveryHandler) StartDelivery(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.StartDelivery(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delivery transit started"})
}

func (h *DeliveryHandler) CompleteDelivery(c *gin.Context) {
	orderID := c.Param("order_id")
	var req requests.OrderStatusChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CompleteDelivery(orderID, req.CourierID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delivery completed successfully"})
}

func (h *DeliveryHandler) PingLocation(c *gin.Context) {
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
