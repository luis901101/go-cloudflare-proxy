package service

import (
	"cloudflare-proxy/conf"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"net/http"
	"time"
)

// BaseService handles image-related routes
type BaseService struct {
	config           conf.Config
	httpClient       *http.Client
	cloudflareClient *cloudflare.Client
}

// NewBaseService creates a new BaseService instance
func NewBaseService(config conf.Config) BaseService {
	return BaseService{
		config:     config,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		cloudflareClient: cloudflare.NewClient(
			option.WithAPIToken(config.Token),
		)}
}

//func (service *BaseService) config() conf.Config {
//	return service._config
//}
//
//func (service *BaseService) httpClient() *http.Client {
//	return service.httpClient
//}
