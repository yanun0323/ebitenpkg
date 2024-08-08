package ebitenpkg

import "sync"

type slices[T any] struct {
	mu    sync.RWMutex
	slice []T
}

func (s *slices[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if index < 0 || index >= len(s.slice) {
		return s.slice[0], false
	}
	return s.slice[index], true
}

func (s *slices[T]) Append(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.slice = append(s.slice, value)
}

func (s *slices[T]) Delete(index int) (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.slice) {
		return s.slice[0], false
	}

	value := s.slice[index]
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return value, true
}

func (s *slices[T]) Range(f func(index int, value T) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, v := range s.slice {
		if !f(i, v) {
			break
		}
	}
}

func (s *slices[T]) FindIndex(find func(T) bool) (index int, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.findIndexWithoutLock(find)
}

func (s *slices[T]) findIndexWithoutLock(find func(T) bool) (index int, ok bool) {
	for i, v := range s.slice {
		if find(v) {
			return i, true
		}
	}

	return 0, false
}

func (s *slices[T]) FindAndSwap(find func(T) bool, value T) (previous T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index, ok := s.findIndexWithoutLock(find)
	if !ok {
		return value, false
	}

	old := s.slice[index]
	s.slice[index] = value
	return old, true
}

func (s *slices[T]) FindAndDelete(find func(T) bool) (value T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index, ok := s.FindIndex(find)
	if !ok {
		return value, false
	}

	old := s.slice[index]
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return old, true
}
