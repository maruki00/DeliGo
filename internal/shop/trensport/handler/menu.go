package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createMenuRequest struct {
	Name string `json:"name" binding:"required"`
}

type addMenuItemRequest struct {
	ProductID string  `json:"product_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

func (h *Handler) CreateMenu(c *gin.Context) {
	shopID := c.Param("shopID")
	var req createMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu, err := h.menuService.CreateMenu(c.Request.Context(), shopID, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

func (h *Handler) GetMenu(c *gin.Context) {
	menuID := c.Param("id")
	menu, err := h.menuService.GetMenu(c.Request.Context(), menuID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (h *Handler) ListMenusByShop(c *gin.Context) {
	shopID := c.Param("shopID")
	menus, err := h.menuService.ListMenusByShop(c.Request.Context(), shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *Handler) AddMenuItem(c *gin.Context) {
	menuID := c.Param("menuID")
	var req addMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu, err := h.menuService.AddItem(c.Request.Context(), menuID, req.ProductID, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (h *Handler) RemoveMenuItem(c *gin.Context) {
	menuID := c.Param("menuID")
	productID := c.Param("productID")
	menu, err := h.menuService.RemoveItem(c.Request.Context(), menuID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (h *Handler) DeleteMenu(c *gin.Context) {
	menuID := c.Param("id")
	if err := h.menuService.DeleteMenu(c.Request.Context(), menuID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
