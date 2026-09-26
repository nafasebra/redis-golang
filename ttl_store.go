package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidExpiration = errors.New("The time is expired")
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
		return ErrEmptyKey
	}

	if !expireAt.After(time.Now()) {
		return ErrInvalidExpiration
	}

	t.data[key] = ttlEntry{
		value:     value,
		expire_at: expireAt,
	}

	return nil
}

func (t *TTLStore) GetWithTTL(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expire_at) {
		delete(t.data, key)
		return "", ErrKeyNotFound
	}

	return entry.value, nil
}

func (t *TTLStore) Delete(key string) error {
	if _, exists := t.data[key]; !exists {
		return ErrKeyNotFound
	}

	delete(t.data, key)
	return nil
}

func (t *TTLStore) Len() int {
	return len(t.data)
}
