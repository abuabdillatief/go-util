package pointer

// FromSliceOrNilIfEmpty returns a pointer to the slice or nil if empty.
func FromSliceOrNilIfEmpty[T any](s []T) *[]T {
	if len(s) == 0 {
		return nil
	}

	return &s
}

// SliceOrNilIfEmpty returns nil if the slice is empty.
func SliceOrNilIfEmpty[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}

	return s
}
