package ebitenpkg

import (
	"sync/atomic"
)

type value[T any] struct {
	value    atomic.Value
	hasValue atomic.Value
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
			return *new(T)
		}

		return value.(T)
	}

	return *new(T)
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
		return *new(T), false
	}

	if swapped == nil {
		return *new(T), false
	}

	return swapped.(T), true
}

func (v *value[T]) Delete() (T, bool) {
	if v.HasValue() {
		v.hasValue.Store(false)
		return v.Load(), true
	}

	return *new(T), false
}

func (v *value[T]) HasValue() bool {
	ok := v.hasValue.Load()
	if ok == nil {
		return false
	}

	return ok.(bool)
}

func (v *value[T]) Copy() value[T] {
	result := value[T]{}
	if v.HasValue() {
		result.Store(v.Load())
	}

	return result
}
