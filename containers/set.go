package containers

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}, size),
	}
}

func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

func (s *Set[T]) Len() int {
	return len(s.data)
}
