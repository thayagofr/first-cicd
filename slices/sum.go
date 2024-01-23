package slices

func Sum[T Number](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}
