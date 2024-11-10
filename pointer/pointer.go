package pointer

// FromVal returns a pointer to the value passed in.
func FromVal[T any](value T) *T {
	return &value
}

// ToVal returns the value at a pointer. If the pointer is nil, a zero value of the type is returned.
func ToVal[T any](pointer *T) T {
	if pointer == nil {
		var value T
		return value
	}

	return *pointer
}

// FromValOrNil returns a pointer if value is not the zero value of the type, otherwise nil
func FromValOrNil[T comparable](value T) *T {
	var zeroValue T
	if value == zeroValue {
		return nil
	}

	return &value
}
