// Package containers provides generic data structures: stack, queue, set and map.
package containers

// Stack is a generic LIFO container.
type Stack[T any] struct {
	data []T
}

// NewStack returns an empty Stack with capacity reserved for size elements.
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, size),
	}
}

// Push adds v to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top element of the stack.
// The second return value is false if the stack is empty.
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

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}
