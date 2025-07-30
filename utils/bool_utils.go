package utils

// BoolPtr converts a bool to a bool pointer
func BoolPtr(s bool) *bool {
	return &s
}

// BoolValue safely converts a bool pointer to bool, returning empty bool if nil
func BoolValue(s *bool) bool {
	return BoolValueOrDefault(s, false)
}

// BoolValueOrDefault safely converts a bool pointer to bool with a default value if nil
func BoolValueOrDefault(s *bool, defaultValue bool) bool {
	if s == nil {
		return defaultValue
	}
	return *s
}
