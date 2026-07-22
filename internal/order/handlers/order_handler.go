package handlers

import (
	"net/http"

	"github.com/maruki00/deligo/internal/order/requests"
	"github.com/maruki00/deligo/internal/order/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req requests.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	order, err := h.svc.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": order})
}

func (h *OrderHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.ConfirmOrder(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Order paid successfully"})
}

func (h *OrderHandler) Accept(c *gin.Context) {
	id := c.Param("id")
	var req requests.AcceptOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	if err := h.svc.AcceptOrder(c.Request.Context(), id, req.Actor); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Order state advanced by actor"})
}

func (h *OrderHandler) Complete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.CompleteDelivery(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delivery completed successfully"})
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.CancelOrder(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Order cancelled successfully"})
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.svc.GetOrder(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": order})
}
