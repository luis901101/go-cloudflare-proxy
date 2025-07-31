package utils

func NullCheck(value ...any) *any {
	if len(value) > 0 {
		for _, v := range value {
			if v != nil {
				return &v
			}
		}
	}
	return nil
}
