package containers

// Queue is a generic FIFO container.
type Queue[T any] struct {
	data []T
}

// NewQueue returns an empty Queue with capacity reserved for size elements.
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, size),
	}
}

// Enqueue adds v to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue removes and returns the front element of the queue.
// The second return value is false if the queue is empty.
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
