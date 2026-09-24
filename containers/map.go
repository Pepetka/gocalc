package containers

// Map is a generic key-value store.
type Map[K comparable, V any] struct {
	data map[K]V
}

// NewMap returns an empty Map with capacity reserved for size entries.
func NewMap[K comparable, V any](size int) *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V, size),
	}
}

// Set stores v under k, replacing any existing value.
func (m *Map[K, V]) Set(k K, v V) {
	m.data[k] = v
}

// Get returns the value stored under k.
// The second return value is false if k is not present in the map.
func (m *Map[K, V]) Get(k K) (V, bool) {
	v, ok := m.data[k]
	return v, ok
}

// Delete removes k and its value from the map.
func (m *Map[K, V]) Delete(k K) {
	delete(m.data, k)
}

// Len returns the number of entries in the map.
func (m *Map[K, V]) Len() int {
	return len(m.data)
}
