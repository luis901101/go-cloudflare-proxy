package controller

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// StreamController handles stream-related routes
type StreamController struct {
	config conf.Config
}

// NewStreamController creates a new StreamController instance
func NewStreamController(config conf.Config) StreamController {
	return StreamController{config: config}
}

// RegisterRoutes registers all stream-related routes
func (controller *StreamController) RegisterRoutes(r *gin.Engine) {
	r.GET("/stream", controller.GetStream)
}

// GetStream handles GET /stream endpoint
func (controller *StreamController) GetStream(c *gin.Context) {
	c.JSON(http.StatusOK, dto.StreamResponseDTO{
		Success: utils.BoolPtr(true),
		Result: &dto.StreamDTO{
			Uid:     utils.StringPtr("asd-qwe-zxc"),
			Created: utils.TimeToUTCTimePtr(time.Now()),
		},
	})
}
