package dto

type GenericResponseDTO[T any] struct {
	Success  *bool                     `json:"success,omitempty"`
	Errors   *[]map[string]interface{} `json:"errors,omitempty"`
	Messages *[]map[string]interface{} `json:"messages,omitempty"`
	Result   *T                        `json:"result,omitempty"`
}
