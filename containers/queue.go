package containers

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, size),
	}
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	l := len(q.data)
	if l == 0 {
		var zero T
		return zero, false
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, true
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}
