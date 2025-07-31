package dto

import "cloudflare-proxy/utils"

type ImageDTO struct {
	ID                *string                  `json:"id,omitempty"`
	Filename          *string                  `json:"filename,omitempty"`
	Meta              **map[string]interface{} `json:"meta,omitempty"`
	Metadata          *map[string]interface{}  `json:"metadata,omitempty"`
	RequireSignedURLs *bool                    `json:"requireSignedURLs,omitempty"`
	Draft             *bool                    `json:"draft,omitempty"`
	Variants          *[]string                `json:"variants,omitempty"`
	Uploaded          *utils.UTCTime           `json:"uploaded,omitempty"`
}

type ImageResponseDTO = GenericResponseDTO[ImageDTO]
