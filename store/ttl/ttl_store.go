package store

import (
	"time"
	"redis-golang/errors"
)

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

func (t *TTLStore) Set(key string, value string, expireAt time.Time) error {
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

func (t *TTLStore) GetWithTTL(key string) (string, error) {
	if key == "" {
		return "", errors.ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expire_at) {
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

    if !exists || !entry.expire_at.After(time.Now()) {
        delete(t.data, oldKey)
        return errors.ErrKeyNotFound
    }

    if existingEntry, exists := t.data[newKey]; exists {
        if existingEntry.expire_at.After(time.Now()) {
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
	return len(t.data)
}
