package kv

import (
	"redis-golang/errors"
	"redis-golang/store"
	"sort"
	"time"
)

var _ store.Storer = (*Store)(nil)

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, error) {
	// simulate time-costing work
	time.Sleep(time.Second)
	
	value, exist := s.data[key]
	if !exist {
		return "", errors.ErrKeyNotFound
	}
	return value, nil
}

func (s *Store) Set(key, value string) error {
	// simulate time-costing work
	time.Sleep(time.Second)

	if key == "" {
		return errors.ErrInvalidKey
	}

	s.data[key] = value

	return nil
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (s *Store) Delete(key string) error {
	if _, exists := s.data[key]; !exists {
		return errors.ErrKeyNotFound
	}

	delete(s.data, key)
	return nil
}

func (s *Store) Len() int {
	return len(s.data)
}

func (s *Store) Clone() *Store {
	cp := &Store{
		data: make(map[string]string, len(s.data)),
	}
	for key, value := range s.data {
		cp.data[key] = value
	}
	return cp
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}
