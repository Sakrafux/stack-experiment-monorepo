package util

func DerefOrDefault[T any](ptr *T, defaultVal T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}
