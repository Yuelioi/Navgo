package cache

import (
	"time"
)

var Cache *Manager

type Manager struct {
	cache map[string]cacheItem
}

func init() {
	Cache = &Manager{
		cache: make(map[string]cacheItem),
	}
}

type cacheItem struct {
	ID          string
	createdTime time.Time
	Content     any
}

func (m *Manager) Add(id string, item any) {
	m.cache[id] = cacheItem{ID: id, createdTime: time.Now(), Content: item}
}

func (m *Manager) Remove(id string) {
	if m.Exists(id) {
		delete(m.cache, id)
	}
}

func (m *Manager) Exists(id string) bool {
	_, ok := m.cache[id]
	return ok
}
