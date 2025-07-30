package dto

type StreamDownloadDTO struct {
	Default *StreamDownloadDefaultDTO `json:"default,omitempty"`
}

type StreamDownloadDefaultDTO struct {
	Status          *string  `json:"status,omitempty"`
	Url             *string  `json:"url,omitempty"`
	PercentComplete *float64 `json:"percentComplete,omitempty"`
}

type StreamDownloadResponseDTO = GenericResponseDTO[StreamDownloadDTO]
