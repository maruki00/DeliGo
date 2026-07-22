package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/restaurant/middlewares"
	"github.com/maruki00/deligo/internal/restaurant/requests"
	"github.com/maruki00/deligo/internal/restaurant/services"
)

type CatalogHandler struct {
	srv services.CatalogService
}

func NewCatalogHandler(srv services.CatalogService) *CatalogHandler {
	return &CatalogHandler{srv: srv}
}

func (h *CatalogHandler) handleError(c *gin.Context, err error) {
	if errors.Is(err, services.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, services.ErrUnauthorized) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not own this restaurant profile"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server boundary error occurred"})
}

func (h *CatalogHandler) CreateRestaurant(c *gin.Context) {
	ownerID := c.MustGet(middlewares.UserContextKey).(string)
	var req requests.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.CreateRestaurant(ownerID, req)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *CatalogHandler) PatchRestaurantStatus(c *gin.Context) {
	ownerID := c.MustGet(middlewares.UserContextKey).(string)
	resID := c.Param("id")

	var req requests.UpdateRestaurantStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ToggleRestaurantStatus(ownerID, resID, *req.IsOpen); err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Restaurant status updated successfully"})
}

func (h *CatalogHandler) AddProduct(c *gin.Context) {
	ownerID := c.MustGet(middlewares.UserContextKey).(string)
	resID := c.Param("id")

	var req requests.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod, err := h.srv.AddProduct(ownerID, resID, req)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, prod)
}

func (h *CatalogHandler) UpdateProduct(c *gin.Context) {
	ownerID := c.MustGet(middlewares.UserContextKey).(string)
	prodID := c.Param("product_id")

	var req requests.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod, err := h.srv.UpdateProduct(ownerID, prodID, req)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, prod)
}

func (h *CatalogHandler) DeleteProduct(c *gin.Context) {
	ownerID := c.MustGet(middlewares.UserContextKey).(string)
	prodID := c.Param("product_id")

	if err := h.srv.DeleteProduct(ownerID, prodID); err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed safely from menu"})
}

func (h *CatalogHandler) GetRestaurantMenu(c *gin.Context) {
	resID := c.Param("id")
	res, err := h.srv.GetMenu(resID)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
