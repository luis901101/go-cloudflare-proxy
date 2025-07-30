package utils

// StringPtr converts a string to a string pointer
func StringPtr(s string) *string {
	return &s
}

// StringValue safely converts a string pointer to string, returning empty string if nil
func StringValue(s *string) string {
	return StringValueOrDefault(s, "")
}

// StringValueOrDefault safely converts a string pointer to string with a default value if nil
func StringValueOrDefault(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return *s
}
