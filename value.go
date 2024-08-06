package ebitenpkg

import "sync/atomic"

type value[T any] struct {
	d T
	v atomic.Value
}

func newValue[T any](d T) value[T] {
	v := atomic.Value{}
	v.Store(d)
	return value[T]{
		d: d,
		v: v,
	}
}

func (v *value[T]) Load() T {
	value := v.v.Load()
	if value == nil {
		return v.d
	}

	return value.(T)
}

func (v *value[T]) Store(d T) {
	v.v.Store(d)
}
