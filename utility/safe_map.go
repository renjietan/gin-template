package utility

import "sync"

type UpdateFunc[T any] func(*T)

type SafeMap[K comparable, V any] struct {
	sync.RWMutex
	data map[K]*V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]*V),
	}
}

func (m *SafeMap[K, V]) SafeUpdate(key K, updates ...UpdateFunc[V]) {
	m.Lock()
	defer m.Unlock()

	item, exists := m.data[key]
	if !exists {
		var zero V
		item = &zero
		m.data[key] = item
	}

	for _, update := range updates {
		update(item)
	}
}

// 返回副本（返回副本，避免外部修改）
func (m *SafeMap[K, V]) GetCopy(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()

	if item, exists := m.data[key]; exists {
		return *item, true
	}
	var zero V
	return zero, false
}

func (m *SafeMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func (m *SafeMap[K, V]) Keys() []K {
	m.RLock()
	defer m.RUnlock()

	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}
