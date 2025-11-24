package handlers

import (
	"net/http"

	"mon-projet/internal/domain"
	"mon-projet/internal/service"

	"github.com/gin-gonic/gin"
)

// ConfigHandler exposes endpoints to inspect and update application settings.
type ConfigHandler struct {
	service *service.ConfigService
}

func NewConfigHandler(s *service.ConfigService) *ConfigHandler {
	return &ConfigHandler{service: s}
}

func (h *ConfigHandler) GetConfiguration(c *gin.Context) {
	cfg, err := h.service.GetConfiguration()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cfg)
}

func (h *ConfigHandler) UpdateConfiguration(c *gin.Context) {
	var payload domain.ApplicationConfig
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateConfiguration(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
