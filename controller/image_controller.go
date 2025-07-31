package controller

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/handler"
	"cloudflare-proxy/service"
	"github.com/gin-gonic/gin"
)

// ImageController handles image-related routes
type ImageController struct {
	config  conf.Config
	service service.ImageService
}

// NewImageController creates a new ImageController instance
func NewImageController(config conf.Config) ImageController {
	return ImageController{
		config:  config,
		service: service.NewImageService(config),
	}
}

// RegisterRoutes registers all image-related routes
func (controller *ImageController) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/image/:id", controller.GetImage)
}

// GetImage handles GET /image endpoint
func (controller *ImageController) GetImage(c *gin.Context) {
	id := c.Params.ByName("id")
	response, status, err := controller.service.GetImage(id)
	handler.HandleResponseWithStatus(c, status, response, err)
}
