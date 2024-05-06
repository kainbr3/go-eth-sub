package storage

import (
	"context"
	"strings"
	"sync"
	"time"

	me "github.com/kainbr3/go-eth-sub/internal/mappedErrors"
)

// the key for map/redis/memcached is composed by:
// transaction/block/hash/from/to
var memoryStorageMap = map[string]any{}

type MemoryStorage struct {
	mu sync.Mutex
}

func NewStorage() (*MemoryStorage, error) {
	if memoryStorageMap == nil {
		memoryStorageMap = map[string]any{}
	}

	return &MemoryStorage{mu: sync.Mutex{}}, nil
}

func (m *MemoryStorage) Get(ctx context.Context, key string) (any, error) {
	addresses := []string{}

	for addr := range memoryStorageMap {
		addresses = append(addresses, strings.Split(addr, "/")[1])
	}

	return addresses, nil
}

func (m *MemoryStorage) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	if _, exists := memoryStorageMap[key]; exists {
		return me.ErrStorageSetDuplicate.BuildErrorMessage(me.STORAGE, nil)
	}

	m.mu.Lock()
	memoryStorageMap[key] = value
	m.mu.Unlock()

	return nil
}

func (m *MemoryStorage) Delete(ctx context.Context, key string) error {
	if _, exists := memoryStorageMap[key]; !exists {
		return me.ErrNotFound.BuildErrorMessage(me.STORAGE, nil)
	}

	m.mu.Lock()
	delete(memoryStorageMap, key)
	m.mu.Unlock()

	return nil
}
