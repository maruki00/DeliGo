package handlers

import (
	"io"
	"net/http"

	"github.com/maruki00/deligo/internal/payment/requests"
	"github.com/maruki00/deligo/internal/payment/service"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	svc           service.PaymentService
	webhookSecret string
}

func NewPaymentHandler(svc service.PaymentService, webhookSecret string) *PaymentHandler {
	return &PaymentHandler{
		svc:           svc,
		webhookSecret: webhookSecret,
	}
}

func (h *PaymentHandler) Charge(c *gin.Context) {
	var req requests.ChargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.svc.ProcessCharge(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": payment})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processing cycle executed", "data": payment})
}

func (h *PaymentHandler) Refund(c *gin.Context) {
	var req requests.RefundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.svc.ProcessRefund(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Refund processed successfully", "data": payment})
}

func (h *PaymentHandler) Webhook(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body payload"})
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")
	err = h.svc.HandleWebhook(c.Request.Context(), payload, sigHeader, h.webhookSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"received": true})
}
