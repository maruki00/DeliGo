package handlers

import (
	"errors"
	"net/http"

	"github.com/maruki00/deligo/internal/user/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.RegisterUser(c.Request.Context(), req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dynamic parameters matching standard UUID format"})
		return
	}

	user, err := h.service.GetUser(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dynamic parameters matching standard UUID format"})
		return
	}

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dynamic parameters matching standard UUID format"})
		return
	}

	if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User sequence successfully removed from core domain"})
}

func (h *UserHandler) Ban(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dynamic parameters matching standard UUID format"})
		return
	}

	var req domain.BanUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.ChangeUserStatus(c.Request.Context(), id, req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrEmailAlreadyExists), errors.Is(err, domain.ErrPhoneAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrInvalidInput):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred internally"})
	}
}
