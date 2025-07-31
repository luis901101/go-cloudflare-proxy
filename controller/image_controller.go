package controller

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/handler"
	"cloudflare-proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (controller *ImageController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/images/:id", controller.get)
	router.POST("/images/direct_upload", controller.directUpload)
}

// get It retrieves an image by its ID from Cloudflare.
func (controller *ImageController) get(c *gin.Context) {
	id := c.Params.ByName("id")
	response, status, err := controller.service.Get(id)
	handler.HandleResponseWithStatus(c, status, response, err)
}

// directUpload generates a direct upload URL for an image.
func (controller *ImageController) directUpload(c *gin.Context) {
	var requestBody *dto.ImageUploadDTO

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handler.HandleResponseWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	response, status, err := controller.service.DirectUpload(requestBody)
	handler.HandleResponseWithStatus(c, status, response, err)
}
