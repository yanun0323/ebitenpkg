package ebitenpkg

import "sync/atomic"

type value[T any] struct {
	defaultValue T
	value        atomic.Value
}

func newValue[T any](d T) value[T] {
	v := atomic.Value{}
	v.Store(d)
	return value[T]{
		defaultValue: d,
		value:        v,
	}
}

func (v *value[T]) Load() T {
	value := v.value.Load()
	if value == nil {
		return v.defaultValue
	}

	return value.(T)
}

func (v *value[T]) Store(d T) {
	v.value.Store(d)
}

func (v *value[T]) Swap(d T) (old T) {
	swapped := v.value.Swap(d)
	if swapped == nil {
		return v.defaultValue
	}

	return swapped.(T)
}
