package dto

import "cloudflare-proxy/utils"

type StreamUploadDTO struct {
	MaxDurationSeconds    *int64         `json:"maxDurationSeconds,omitempty"`
	AllowedOrigins        *[]string      `json:"allowedOrigins,omitempty"`
	Creator               *string        `json:"creator,omitempty"`
	Expiry                *utils.UTCTime `json:"expiry,omitempty"`
	Meta                  *any           `json:"meta,omitempty"`
	RequireSignedURLs     *bool          `json:"requireSignedURLs,omitempty"`
	ScheduledDeletion     *utils.UTCTime `json:"scheduledDeletion,omitempty"`
	ThumbnailTimestampPct *float64       `json:"thumbnailTimestampPct,omitempty"`
	Watermark             *Watermark     `json:"watermark,omitempty"`
	UploadCreator         *string        `json:"uploadCreator,omitempty"`
}
