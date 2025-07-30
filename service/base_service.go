package service

import (
	"cloudflare-proxy/conf"
	"net/http"
	"time"
)

// BaseService handles image-related routes
type BaseService struct {
	config conf.Config
	client *http.Client
}

// NewBaseService creates a new BaseService instance
func NewBaseServiceService(config conf.Config) BaseService {
	return BaseService{config: config, client: &http.Client{Timeout: 30 * time.Second}}
}

//func (service *BaseService) config() conf.Config {
//	return service._config
//}
//
//func (service *BaseService) client() *http.Client {
//	return service.client
//}
