package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	shared_valueobject "github.com/maruki00/deligo/internal/shared/value_object"
)

func (h *Handler) CreateProfile(c *gin.Context) {
	var payload struct {
		UserID   string `json:"user_id" binding:"required"`
		FullName string `json:"full_name" binding:"required"`
		Avatar   string `json:"avatar"`
		Bio      string `json:"bio"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.app.ProfileRepo.Save(c.Request.Context(), &model.Profile{
		ID:       shared_valueobject.NewID(),
		UserID:   shared_valueobject.ID(payload.UserID),
		FullName: payload.FullName,
		Avatar:   payload.Avatar,
		Bio:      payload.Bio,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "profile created"})
}

func (h *Handler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.app.ProfileRepo.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         string(profile.GetID()),
		"user_id":    string(profile.GetUserID()),
		"full_name":  profile.GetFullName(),
		"avatar":     profile.GetAvatar(),
		"bio":        profile.GetBio(),
		"created_at": profile.GetCreatedAt(),
		"updated_at": profile.GetUpdatedAt(),
	})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.app.ProfileRepo.Update(c.Request.Context(), id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated"})
}

func (h *Handler) UpdateAvatar(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		Avatar string `json:"avatar" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.app.ProfileRepo.UpdateAvatar(c.Request.Context(), id, payload.Avatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "avatar updated"})
}

func (h *Handler) DisableProfile(c *gin.Context) {
	id := c.Param("id")
	err := h.app.ProfileRepo.Disable(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile disabled"})
}
