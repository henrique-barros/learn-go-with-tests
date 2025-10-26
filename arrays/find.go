package arrays

func Find[T any](values []T, fn func(val T) bool) (bool, T) {
	var zero T
	for _, cur := range values {
		if fn(cur) {
			return true, cur
		}
	}
	return false, zero
}
