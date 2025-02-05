package utils

func TakeFirst[T any](defaultData T, options []T) T {
	if len(options) >= 1 {
		return options[0]
	}
	return defaultData
}
