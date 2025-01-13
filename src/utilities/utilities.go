package util

func IfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func IfZero(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}
