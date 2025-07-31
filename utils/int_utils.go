package utils

// Int64Ptr converts a int64 to a int64 pointer
func Int64Ptr(s int64) *int64 {
	return &s
}

// Int64Value safely converts a int64 pointer to int64, returning empty int64 if nil
func Int64Value(s *int64) int64 {
	return Int64ValueOrDefault(s, 0)
}

// Int64ValueOrDefault safely converts a int64 pointer to int64 with a default value if nil
func Int64ValueOrDefault(s *int64, defaultValue int64) int64 {
	if s == nil {
		return defaultValue
	}
	return *s
}
