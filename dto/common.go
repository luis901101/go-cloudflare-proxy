package dto

import "time"

type Input struct {
	Width  *int64 `json:"width,omitempty"`
	Height *int64 `json:"height,omitempty"`
}

type Watermark struct {
	DownloadedFrom *string    `json:"downloadedFrom,omitempty"`
	Padding        *float64   `json:"padding,omitempty"`
	Uid            *string    `json:"uid,omitempty"`
	Size           *int64     `json:"size,omitempty"`
	Created        *time.Time `json:"created,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Width          *int64     `json:"width,omitempty"`
	Scale          *float64   `json:"scale,omitempty"`
	Position       *string    `json:"position,omitempty"`
	Opacity        *float64   `json:"opacity,omitempty"`
	Height         *int64     `json:"height,omitempty"`
}
