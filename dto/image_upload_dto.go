package dto

type ImageUploadDTO struct {
	ID                *string `json:"id,omitempty"`
	Creator           *string `json:"creator,omitempty"`
	File              *string `json:"file,omitempty"`
	Metadata          *any    `json:"metadata,omitempty"`
	RequireSignedURLs *bool   `json:"requireSignedURLs,omitempty"`
	Url               *string `json:"url,omitempty"`
}
