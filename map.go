package ebitenpkg

import "sync"

type maps[K comparable, V any] struct {
	data sync.Map
}

func (m *maps[K, V]) Get(key K) (V, bool) {
	v, ok := m.data.Load(key)
	if !ok {
		return v.(V), false
	}
	return v.(V), true
}

func (m *maps[K, V]) Set(key K, value V) {
	m.data.Store(key, value)
}

func (m *maps[K, V]) Delete(key K) (value V, loaded bool) {
	val, ok := m.data.LoadAndDelete(key)
	return val.(V), ok
}

func (m *maps[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *maps[K, V]) Swap(key K, value V) (previous V, ok bool) {
	old, ok := m.data.Swap(key, value)
	if ok {
		return old.(V), true
	}

	var zero V
	return zero, false
}
