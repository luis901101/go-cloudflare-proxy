package dto

import (
	"cloudflare-proxy/utils"
)

type StreamDTO struct {
	Preview               *string                 `json:"preview,omitempty"`
	UploadExpiry          *utils.UTCTime          `json:"uploadExpiry,omitempty"`
	Creator               *string                 `json:"creator,omitempty"`
	Thumbnail             *string                 `json:"thumbnail,omitempty"`
	RequireSignedURLs     *bool                   `json:"requireSignedURLs,omitempty"`
	Watermark             *Watermark              `json:"watermark,omitempty"`
	Created               *utils.UTCTime          `json:"created,omitempty"`
	ReadyToStreamAt       *utils.UTCTime          `json:"readyToStreamAt,omitempty"`
	ThumbnailTimestampPct *float64                `json:"thumbnailTimestampPct,omitempty"`
	ReadyToStream         *bool                   `json:"readyToStream,omitempty"`
	Duration              *int64                  `json:"duration,omitempty"`
	Input                 *Input                  `json:"input,omitempty"`
	Uid                   *string                 `json:"uid,omitempty"`
	AllowedOrigins        *[]string               `json:"allowedOrigins,omitempty"`
	Size                  *int64                  `json:"size,omitempty"`
	LiveInput             *string                 `json:"liveInput,omitempty"`
	Meta                  *map[string]interface{} `json:"meta,omitempty"`
	Uploaded              *utils.UTCTime          `json:"uploaded,omitempty"`
	Modified              *utils.UTCTime          `json:"modified,omitempty"`
	Playback              *StreamPlayback         `json:"playback,omitempty"`
	ScheduledDeletion     *utils.UTCTime          `json:"scheduledDeletion,omitempty"`
	MaxDurationSeconds    *int64                  `json:"maxDurationSeconds,omitempty"`
	Status                *StreamStatus           `json:"status,omitempty"`
}

type StreamResponseDTO = GenericResponseDTO[StreamDTO]

type StreamPlayback struct {
	Dash *string `json:"dash,omitempty"`
	HLS  *string `json:"hls,omitempty"`
}

type StreamStatus struct {
	ErrorReasonCode *string `json:"errorReasonCode,omitempty"`
	ErrorReasonText *string `json:"errorReasonText,omitempty"`
	State           *string `json:"state,omitempty"`
	PctComplete     *string `json:"pctComplete,omitempty"`
}
