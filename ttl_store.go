package main

import (
	"fmt"
	"time"
	"errors"
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

func (t *TTLStore) Set(key string, value string, ttl time.Time) error {
	if key == "" {
		return ErrEmptyKey
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
