package containers

type Map[K comparable, V any] struct {
	data map[K]V
}

func NewMap[K comparable, V any](size int) *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V, size),
	}
}

func (m *Map[K, V]) Set(k K, v V) {
	m.data[k] = v
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	v, ok := m.data[k]
	return v, ok
}

func (m *Map[K, V]) Delete(k K) {
	delete(m.data, k)
}

func (m *Map[K, V]) Len() int {
	return len(m.data)
}
