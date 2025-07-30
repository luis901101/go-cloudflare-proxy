package service

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/handler"
	"encoding/json"
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
func (service ImageService) GetImage(id string) (result dto.ImageResponseDTO, status int, err error) {
	// Build the URL using config methods
	url := fmt.Sprintf("%s/%s", service.config.ImageUrl(), id)

	// Create
	request, status, err := handler.NewRequest(http.MethodGet, url, service.config, nil)
	if err != nil {
		return result, status, err
	}

	// Handle request
	response, status, err := handler.HandleRequest(service.client, request)
	if status != http.StatusOK || err != nil {
		return result, status, err
	}
	//goland:noinspection ALL
	defer response.Body.Close()

	// Decode JSON body
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("failed to decode result: %w", err)
	}

	return result, http.StatusOK, nil
}
