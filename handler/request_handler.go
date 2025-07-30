package handler

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/utils"
	"fmt"
	"io"
	"net/http"
)

func NewRequest(method, url string, config conf.Config, body io.Reader) (*http.Request, int, error) {
	// Create HTTP request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to create request: %w", err)
	}
	// Set authorization header using config method
	req.Header.Set(utils.HeaderAuthorizationKey, config.Authorization())
	req.Header.Set(utils.HeaderContentTypeKey, utils.HeaderApplicationJSONValue)
	return req, http.StatusOK, nil
}

func HandleRequest(client *http.Client, request *http.Request) (*http.Response, int, error) {
	// Make the HTTP request
	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to make request: %w", err)
	}

	// Check status code
	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("API request failed with status %d", response.StatusCode)
	}

	return response, http.StatusOK, nil
}
