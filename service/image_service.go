package service

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/handler"
	"fmt"
	"net/http"
)

// ImageService handles Cloudflare image related operations
// Official documentation: https://developers.cloudflare.com/images/
type ImageService struct {
	BaseService
}

// NewImageService creates a new ImageService instance
func NewImageService(config conf.Config) ImageService {
	return ImageService{NewBaseServiceService(config)}
}

// GetImage handles GET image by id from Cloudflare
func (service ImageService) GetImage(id string) (result *any /*result *dto.ImageResponseDTO*/, status int, err error) {
	// Build the URL using config methods
	url := fmt.Sprintf("%s/%s", service.config.ImageUrl(), id)

	// Create
	request, status, err := handler.NewRequest(http.MethodGet, url, service.config, nil)
	if err != nil {
		return result, status, err
	}

	return handler.HandleRequestWithResult[any /*dto.ImageResponseDTO*/](service.client, request)
}
