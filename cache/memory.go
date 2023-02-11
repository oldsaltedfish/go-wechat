package cache

import (
	"context"
	"sync"
)

type Memory struct {
	store *sync.Map
}

func NewMemory() *Memory {
	m := &Memory{
		store: &sync.Map{},
	}
	return m
}

func (m *Memory) Save(ctx context.Context, key string, data any) error {
	m.store.Store(key, data)
	return nil
}

func (m *Memory) Get(ctx context.Context, key string, data any) error {
	data, ok := m.store.Load(key)
	if ok {
		return nil
	}
	return ErrCacheNil
}
