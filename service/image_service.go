package service

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/handler"
	"cloudflare-proxy/utils"
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"net/http"
)

// ImageService handles Cloudflare image related operations
// Official documentation: https://developers.cloudflare.com/images/
type ImageService struct {
	BaseService
}

// NewImageService creates a new ImageService instance
func NewImageService(config conf.Config) ImageService {
	return ImageService{NewBaseService(config)}
}

// GetManually handles GET image by id from Cloudflare
func (service ImageService) GetManually(id string) (result *any /*result *dto.ImageResponseDTO*/, status int, err error) {
	// Build the URL using config methods
	url := fmt.Sprintf("%s/%s", service.config.ImageUrl(), id)

	// Create
	request, status, err := handler.NewRequest(http.MethodGet, url, service.config, nil)
	if err != nil {
		return result, status, err
	}

	return handler.HandleRequestWithResult[any /*dto.ImageResponseDTO*/](service.httpClient, request)
}

// Get handles GET image by id from Cloudflare
func (service ImageService) Get(id string) (result any /*result *dto.ImageResponseDTO*/, status int, err error) {

	response, err := service.cloudflareClient.Images.V1.Get(
		context.TODO(),
		id,
		images.V1GetParams{
			AccountID: cloudflare.F(service.config.Account),
		},
	)
	return response, http.StatusOK, err
}

// DirectUpload handles creating a direct upload for an Image in Cloudflare
// Official documentation: https://developers.cloudflare.com/api/resources/images/subresources/v2/subresources/direct_uploads/methods/create/
func (service ImageService) DirectUpload(requestBody *dto.ImageUploadDTO) (result any /*result *dto.ImageResponseDTO*/, status int, err error) {

	var params = images.V2DirectUploadNewParams{
		AccountID: cloudflare.F(service.config.Account),
	}

	if requestBody.RequireSignedURLs != nil {
		params.RequireSignedURLs = cloudflare.F(utils.BoolValue(requestBody.RequireSignedURLs))
	}
	if requestBody.ID != nil {
		params.ID = cloudflare.F[string](utils.StringValue(requestBody.ID))
	}
	//if requestBody.Metadata != nil {
	//	params.Metadata = cloudflare.F(*requestBody.Metadata)
	//}

	response, err := service.cloudflareClient.Images.V2.DirectUploads.New(
		context.TODO(),
		params,
	)

	return response, http.StatusOK, err
}
