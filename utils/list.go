package utils

func Prepend[T any](slice []T, elements ...T) []T {
	return append(elements, slice...)
}
