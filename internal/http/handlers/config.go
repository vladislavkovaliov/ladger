package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladislavkovaliov/ledger/internal/config"
)

// ConfigHandler godoc
// @Summary Get config
// @Description Check service config
// @Tags system
// @Produce json
// @Success 200 {object} dto.ConfigResponse
// @Router /config [get]
func ConfigHandler(c *gin.Context, cfg *config.Config) {
	c.JSON(http.StatusOK, cfg)
}
