package memory

import (
	"sync"

	"github.com/khan745/gokvdb/internal/pkg/storage"
)

//Storage represent enthity that store values by key in memory. Implements Storage interface
type Storage struct {
	mu    sync.RWMutex
	keys  map[string]storage.Key
	items map[storage.Key]storage.Value
}

//New is constructor for new storage
func New() *Storage {
	return &Storage{
		keys:  make(map[string]storage.Key),
		items: make(map[storage.Key]storage.Value),
	}
}

//Put put value to storage by key
func (strg *Storage) Put(key storage.Key, value storage.Value) error {
	strg.mu.Lock()
	strg.keys[key.Val()] = key
	strg.items[key] = value
	strg.mu.Unlock()
	return nil
}

//Get get value of storage by key
func (strg *Storage) Get(key storage.Key) (storage.Value, error) {
	strg.mu.RLock()
	val := strg.items[key]
	strg.mu.RUnlock()
	return val, nil
}

//Del delete key and value
func (strg *Storage) Del(key storage.Key) error {
	strg.mu.Lock()
	delete(strg.keys, key.Val())
	delete(strg.items, key)
	strg.mu.Unlock()
	return nil
}

//GetKey get key by name
func (strg *Storage) GetKey(keyName string) (storage.Key, error) {
	strg.mu.RLock()
	key, ok := strg.keys[keyName]
	strg.mu.RUnlock()
	if !ok {
		return storage.Key{}, storage.ErrKeyNotExists
	}
	return key, nil
}

//Keys return all stored keys
func (strg *Storage) Keys() ([]storage.Key, error) {
	strg.mu.RLock()
	keys := make([]storage.Key, 0, len(strg.keys))
	for _, key := range strg.keys {
		keys = append(keys, key)
	}
	strg.mu.RUnlock()
	return keys, nil
}
