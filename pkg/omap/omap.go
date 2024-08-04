package omap

import (
	"sync"
)

type OrderedMap[K comparable, V any] struct {
	keys   []K
	values map[K]V
	mu     *sync.RWMutex
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
		mu:     &sync.RWMutex{},
	}
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.values[key]
	if !ok {
		m.keys = append(m.keys, key)
	}

	m.values[key] = value
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.values[key]

	return value, ok
}

func (m *OrderedMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.values[key]
	if ok {
		delete(m.values, key)

		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

func (m *OrderedMap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return append([]K{}, m.keys...)
}
