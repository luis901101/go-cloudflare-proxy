package handler

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/utils"
	"encoding/json"
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

// HandleRequestWithResponse Handles the HTTP request and returns the response along with status code and error if any
// It's your responsibility to Close the response body after using it
func HandleRequestWithResponse(client *http.Client, request *http.Request) (*http.Response, int, error) {
	// Make the HTTP request
	response, err := client.Do(request)
	if err != nil {
		return response, http.StatusInternalServerError, fmt.Errorf("failed to make request: %w", err)
	}

	// Check status code
	if response.StatusCode != http.StatusOK {
		return response, response.StatusCode, fmt.Errorf("API request failed with status %d", response.StatusCode)
	}

	return response, http.StatusOK, nil
}

// HandleRequestWithResult handles the HTTP request and decodes the JSON response into the provided result type.
// It returns the result, status code, and error if any.
func HandleRequestWithResult[T any](client *http.Client, request *http.Request) (result *T, status int, err error) {
	// Handle request
	response, status, err := HandleRequestWithResponse(client, request)
	//goland:noinspection ALL
	defer response.Body.Close()

	errDecoding := json.NewDecoder(response.Body).Decode(&result)

	if status != http.StatusOK || err != nil {
		return result, status, err
	}

	// Decode JSON body
	if errDecoding != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("failed to decode result: %w", errDecoding)
	}

	return result, http.StatusOK, nil
}
