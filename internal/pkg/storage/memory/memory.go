package memory

import (
	"sync"

	"github.com/khan745/gokvdb/internal/pkg/storage"
)

//Storage represent enthity that store values by key in memory. Implements Storage interface
type Storage struct {
	mu    sync.RWMutex
	items map[storage.Key]*storage.Value
}

//New is constructor for new storage
func New() *Storage {
	return &Storage{
		items: make(map[storage.Key]*storage.Value),
	}
}

//Put put value to storage by key
func (strg *Storage) Put(key storage.Key, setter storage.ValueSetter) error {
	strg.mu.Lock()

	value := strg.items[key]

	var err error

	if value, err = setter(value); err == nil {
		if value == nil {
			delete(strg.items, key)
		} else {
			strg.items[key] = value
		}
	}

	strg.mu.Unlock()
	return err
}

//Get get value of storage by key
func (strg *Storage) Get(key storage.Key) (*storage.Value, error) {
	strg.mu.RLock()
	defer strg.mu.RUnlock()

	if value, exist := strg.items[key]; exist {
		return value, nil
	}
	return nil, storage.ErrKeyNotExists
}

//Del delete key and value
func (strg *Storage) Del(key storage.Key) error {
	strg.mu.Lock()
	delete(strg.items, key)
	strg.mu.Unlock()
	return nil
}

//Keys return all stored keys
func (strg *Storage) Keys() ([]storage.Key, error) {
	strg.mu.RLock()
	keys := make([]storage.Key, 0, len(strg.items))
	for key := range strg.items {
		keys = append(keys, key)
	}
	strg.mu.RUnlock()
	return keys, nil
}
