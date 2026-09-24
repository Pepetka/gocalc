// Package containers provides stack data structure.
package containers

type Stack[T any] struct {
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, size),
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	l := len(s.data)
	if l == 0 {
		var zero T
		return zero, false
	}
	v := s.data[l-1]
	s.data = s.data[:l-1]
	return v, true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}
