package util

func IfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}