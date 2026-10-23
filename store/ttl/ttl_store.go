package ttl

import (
	"redis-golang/errors"
	"redis-golang/store"
	"sort"
	"time"
)

var _ store.Storer = (*TTLStore)(nil)

type TTLStore struct {
	data map[string]ttlEntry
}

type ttlEntry struct {
	value     string
	expire_at time.Time
}

func NewTTLStore() *TTLStore {
	return &TTLStore{
		data: make(map[string]ttlEntry),
	}
}

func (t *TTLStore) SetWithTTL(key string, value string, expireAt time.Time) error {
	if key == "" {
		return errors.ErrEmptyKey
	}

	if !expireAt.After(time.Now()) {
		return errors.ErrInvalidExpiration
	}

	t.data[key] = ttlEntry{
		value:     value,
		expire_at: expireAt,
	}

	return nil
}

func (t *TTLStore) GetWithTTL(key string) (string, error) { return t.Get(key) }

func (t *TTLStore) Get(key string) (string, error) {
	if key == "" {
		return "", errors.ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || entry.expired() {
		delete(t.data, key)
		return "", errors.ErrKeyNotFound
	}

	return entry.value, nil
}

func (t *TTLStore) Delete(key string) error {
	if _, exists := t.data[key]; !exists {
		return errors.ErrKeyNotFound
	}

	delete(t.data, key)
	return nil
}

func (t *TTLStore) Rename(oldKey, newKey string) error {
	if oldKey == "" || newKey == "" {
		return errors.ErrInvalidKey
	}

	entry, exists := t.data[oldKey]

	if !exists || entry.expired() {
		delete(t.data, oldKey)
		return errors.ErrKeyNotFound
	}

	if existingEntry, exists := t.data[newKey]; exists {
		if !existingEntry.expired() {
			return errors.ErrKeyAlreadyExists
		}

		delete(t.data, newKey)
	}

	t.data[newKey] = entry
	delete(t.data, oldKey)

	return nil
}

// func (t *TTLStore) Exist(key string) (bool, error) {
// 	if key == "" {
// 		return false, ErrEmptyKey
// 	}

// 	return false, nil
// }

// func (t *TTLStore) TTL(time time.Time) error {
// 	return nil
// }

// func (t *TTLStore) Refresh(value string) error {}

// func (t *TTLStore) Keys() error {
// 	return nil
// }

func (t *TTLStore) Len() int {
	return len(t.Keys())
}

// Set stores a value without expiration.
func (t *TTLStore) Set(key, value string) error {
	if key == "" {
		return errors.ErrEmptyKey
	}
	t.data[key] = ttlEntry{value: value}
	return nil
}

func (entry ttlEntry) expired() bool {
	return !entry.expire_at.IsZero() && !entry.expire_at.After(time.Now())
}

func (t *TTLStore) Keys() []string {
	keys := make([]string, 0, len(t.data))
	for key, entry := range t.data {
		if entry.expired() {
			delete(t.data, key)
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
