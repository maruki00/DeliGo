package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createShopRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	OwnerID string `json:"owner_id"`
}

type updateShopRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (h *Handler) CreateShop(c *gin.Context) {
	var req createShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop, err := h.shopService.CreateShop(c.Request.Context(), req.Name, req.Address, req.Phone, req.OwnerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shop)
}

func (h *Handler) GetShop(c *gin.Context) {
	shopID := c.Param("id")
	shop, err := h.shopService.GetShop(c.Request.Context(), shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shop)
}

func (h *Handler) ListShops(c *gin.Context) {
	shops, err := h.shopService.ListShops(c.Request.Context(), 20, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shops)
}

func (h *Handler) UpdateShop(c *gin.Context) {
	shopID := c.Param("id")
	var req updateShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop, err := h.shopService.UpdateShop(c.Request.Context(), shopID, req.Name, req.Address, req.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shop)
}

func (h *Handler) DeleteShop(c *gin.Context) {
	shopID := c.Param("id")
	if err := h.shopService.DeleteShop(c.Request.Context(), shopID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
