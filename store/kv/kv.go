package store

import (
	"sort"
	"redis-golang/errors"
)

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	value, exist := s.data[key]
	return value, exist
}

func (s *Store) Set(key, value string) error {
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

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}
