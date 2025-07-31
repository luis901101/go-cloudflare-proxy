package controller

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/handler"
	"cloudflare-proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StreamController handles stream-related routes
type StreamController struct {
	config  conf.Config
	service service.StreamService
}

// NewStreamController creates a new StreamController instance
func NewStreamController(config conf.Config) StreamController {
	return StreamController{
		config:  config,
		service: service.NewStreamService(config),
	}
}

// RegisterRoutes registers all stream-related routes
func (controller *StreamController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/stream/:id", controller.get)
	router.POST("/stream/direct_upload", controller.directUpload)
}

// get handles GET /stream endpoint
func (controller *StreamController) get(c *gin.Context) {
	id := c.Params.ByName("id")
	response, status, err := controller.service.Get(id)
	handler.HandleResponseWithStatus(c, status, response, err)
}

// directUpload generates a direct upload URL for a stream.
func (controller *StreamController) directUpload(c *gin.Context) {
	var requestBody *dto.StreamUploadDTO

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handler.HandleResponseWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	response, status, err := controller.service.DirectUpload(requestBody)
	handler.HandleResponseWithStatus(c, status, response, err)
}
