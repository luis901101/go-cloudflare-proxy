package utils

import (
	"fmt"
	"strconv"
	"time"
)

// TimePtr converts a time.Time to a time.Time pointer
func TimePtr(s time.Time) *time.Time {
	return &s
}

// TimeValue safely converts a time.Time pointer to time.Time, returning empty time.Time if nil
func TimeValue(s *time.Time) time.Time {
	return TimeValueOrDefault(s, time.Now())
}

// TimeValueOrDefault safely converts a time.Time pointer to time.Time with a default value if nil
func TimeValueOrDefault(s *time.Time, defaultValue time.Time) time.Time {
	if s == nil {
		return defaultValue
	}
	return *s
}

// TimeToUTCTimePtr converts a time.Time to a time.UTCTime pointer
func TimeToUTCTimePtr(s time.Time) *UTCTime {
	return UTCTimePtr(UTCTime(s.UTC()))
}

// UTCTimePtr converts a time.UTCTime to a time.UTCTime pointer
func UTCTimePtr(s UTCTime) *UTCTime {
	return &s
}

// UTCTimeValue safely converts a UTCTime pointer to UTCTime, returning empty UTCTime if nil
func UTCTimeValue(s *UTCTime) UTCTime {
	return UTCTimeValueOrDefault(s, UTCTime(time.Now()))
}

// UTCTimeValueOrDefault safely converts a UTCTime pointer to UTCTime with a default value if nil
func UTCTimeValueOrDefault(s *UTCTime, defaultValue UTCTime) UTCTime {
	if s == nil {
		return defaultValue
	}
	return *s
}

// UTCTime is a custom time type that ensures time is always handled
// in UTC and marshaled to JSON in the standard RFC3339 format (Zulu time).
type UTCTime time.Time

// MarshalJSON implements the json.Marshaler interface.
// It formats the time as a UTC string in RFC3339 format.
func (t *UTCTime) MarshalJSON() ([]byte, error) {
	// Convert our custom type to a standard time.Time object
	if t == nil {
		return []byte(nil), nil
	}

	tValue := UTCTimeValue(t)
	stdTime := time.Time(tValue)

	// Format the time in UTC. The 'Z' at the end stands for Zulu (UTC).
	// The result is quoted to be a valid JSON string.
	formatted := fmt.Sprintf("\"%s\"", stdTime.UTC().Format(time.RFC3339))
	return []byte(formatted), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It parses a string from JSON and ensures the resulting time is in UTC.
func (t *UTCTime) UnmarshalJSON(b []byte) error {
	// First, unquote the JSON string.
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	// Parse the time string using the RFC3339 layout.
	// This layout is flexible and handles various ISO 8601 formats.
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	// Convert the parsed time to UTC and assign it to our custom type.
	*t = UTCTime(parsedTime.UTC())
	return nil
}
