package service

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/dto"
	"cloudflare-proxy/handler"
	"cloudflare-proxy/utils"
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/stream"
	"net/http"
	"time"
)

// StreamService handles Cloudflare stream related operations
// Official documentation: https://developers.cloudflare.com/stream
type StreamService struct {
	BaseService
}

// NewStreamService creates a new StreamService instance
func NewStreamService(config conf.Config) StreamService {
	return StreamService{NewBaseService(config)}
}

// GetManually handles GET stream by id from Cloudflare
func (service StreamService) GetManually(id string) (result *any /*result *dto.StreamResponseDTO*/, status int, err error) {
	// Build the URL using config methods
	url := fmt.Sprintf("%s/%s", service.config.StreamUrl(), id)

	// Create
	request, status, err := handler.NewRequest(http.MethodGet, url, service.config, nil)
	if err != nil {
		return result, status, err
	}

	return handler.HandleRequestWithResult[any /*dto.StreamResponseDTO*/](service.httpClient, request)
}

// Get handles GET stream by id from Cloudflare
func (service StreamService) Get(id string) (result any /*result *dto.StreamResponseDTO*/, status int, err error) {

	response, err := service.cloudflareClient.Stream.Get(
		context.TODO(),
		id,
		stream.StreamGetParams{
			AccountID: cloudflare.F(service.config.Account),
		},
	)
	return response, http.StatusOK, err
}

// DirectUpload handles creating a direct upload for a Stream in Cloudflare
// Official documentation: https://developers.cloudflare.com/api/resources/streams/subresources/v2/subresources/direct_uploads/methods/create/
func (service StreamService) DirectUpload(requestBody *dto.StreamUploadDTO) (result any /*result *dto.StreamResponseDTO*/, status int, err error) {

	var params = stream.DirectUploadNewParams{
		AccountID: cloudflare.F(service.config.Account),
	}

	if requestBody.MaxDurationSeconds != nil {
		params.MaxDurationSeconds = cloudflare.F(*requestBody.MaxDurationSeconds)
	}
	if requestBody.AllowedOrigins != nil {
		params.AllowedOrigins = cloudflare.F(*requestBody.AllowedOrigins)
	}
	if requestBody.Creator != nil {
		params.Creator = cloudflare.F(*requestBody.Creator)
	}
	if requestBody.Expiry != nil {
		params.Expiry = cloudflare.F(time.Time(utils.UTCTimeValue(requestBody.Expiry)))
	}
	if requestBody.Meta != nil {
		params.Meta = cloudflare.F(*requestBody.Meta)
	}
	if requestBody.RequireSignedURLs != nil {
		params.RequireSignedURLs = cloudflare.F(*requestBody.RequireSignedURLs)
	}
	if requestBody.ScheduledDeletion != nil {
		params.ScheduledDeletion = cloudflare.F(time.Time(utils.UTCTimeValue(requestBody.ScheduledDeletion)))
	} else {
		params.ScheduledDeletion = cloudflare.Null[time.Time]()
	}
	if requestBody.ThumbnailTimestampPct != nil {
		params.ThumbnailTimestampPct = cloudflare.F(*requestBody.ThumbnailTimestampPct)
	}
	if requestBody.Watermark != nil && requestBody.Watermark.Uid != nil {
		params.Watermark = cloudflare.F(stream.DirectUploadNewParamsWatermark{
			UID: cloudflare.F(*requestBody.Watermark.Uid),
		})
	} else {
		params.Watermark = cloudflare.Null[stream.DirectUploadNewParamsWatermark]()
	}
	if requestBody.UploadCreator != nil {
		params.UploadCreator = cloudflare.F(*requestBody.UploadCreator)
	}

	response, err := service.cloudflareClient.Stream.DirectUpload.New(
		context.TODO(),
		params,
	)

	return response, http.StatusOK, err
}
