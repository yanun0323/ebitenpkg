package ebitenpkg

import (
	"sync/atomic"
)

type value[T any] struct {
	defaultValue T
	value        atomic.Value
	hasValue     atomic.Value
}

// newValue panics if val is nil
func newValue[T any](val ...T) value[T] {
	result := value[T]{}

	if len(val) != 0 {
		result.Store(val[0])
	}

	return result
}

func (v *value[T]) Load() T {
	if v.HasValue() {
		value := v.value.Load()
		if value == nil {
			return v.defaultValue
		}

		return value.(T)
	}

	return v.defaultValue
}

// Store panics if val is nil
func (v *value[T]) Store(val T) {
	v.value.Store(val)
	v.hasValue.Store(true)
}

// Swap panics if val is nil
func (v *value[T]) Swap(val T) (T, bool) {
	swapped := v.value.Swap(val)
	if !v.HasValue() {
		v.hasValue.Store(true)
		return v.defaultValue, false
	}

	if swapped == nil {
		return v.defaultValue, false
	}

	return swapped.(T), true
}

func (v *value[T]) Delete() (T, bool) {
	if v.HasValue() {
		v.hasValue.Store(false)
		return v.Load(), true
	}

	return v.defaultValue, false
}

func (v *value[T]) HasValue() bool {
	ok := v.hasValue.Load()
	if ok == nil {
		return false
	}

	return ok.(bool)
}
