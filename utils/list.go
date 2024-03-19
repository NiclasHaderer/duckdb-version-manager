package utils

func Prepend[T any](slice []T, elements ...T) []T {
	return append(elements, slice...)
}

func Map[T any, R any](elements []T, f func(element T) R) []R {
	var result []R
	for _, element := range elements {
		result = append(result, f(element))
	}
	return result
}
