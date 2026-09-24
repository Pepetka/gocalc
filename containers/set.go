package containers

// Set is a generic collection of unique values.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet returns an empty Set with capacity reserved for size elements.
func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}, size),
	}
}

// Add inserts v into the set.
func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

// Has reports whether v is present in the set.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

// Remove deletes v from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.data)
}
