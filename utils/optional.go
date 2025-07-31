package utils

func NullCheck(value ...any) *any {
	if value != nil {
		for _, v := range value {
			if v != nil {
				return &v
			}
		}
	}
	return nil
}

//func NullCheck2[T any](value T, other T) *T {
//	if value == nil {
//		return &other
//	}
//	return &value
//}
