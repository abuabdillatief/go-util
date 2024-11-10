package iterator

// ForEach iterates over a slice and applies a function to each element.
func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

// Map iterates over a slice and applies a function to each element, returning a new slice with the results.
func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	
	return result
}

// Filter iterates over a slice and applies a function to each element, returning a new slice with the elements that pass the filter.
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

// Reduce iterates over a slice and applies a function to each element, returning a single value.
func Reduce[T any, R any](slice []T, fn func(R, T) R, initial R) R {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}

	return result
}
