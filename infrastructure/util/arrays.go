package util

func Contains[T comparable](array []T, element T) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
