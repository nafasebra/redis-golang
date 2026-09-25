package main

import (
	"errors"
	"sort"
)

var (
	ErrKeyNotFound      = errors.New("key not found")
	ErrKeyAlreadyExists = errors.New("key already exists")
	ErrInvalidKey        = errors.New("key cannot be empty")
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
		return ErrInvalidKey
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) error {
	if _, exists := s.data[key]; !exists {
		return ErrKeyNotFound
	}

	delete(s.data, key)
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

func (s *Store) Rename(oldKey, newKey string) error {
	if oldKey == "" || newKey == "" {
		return ErrInvalidKey
	}

	value, exists := s.data[oldKey]

	if !exists {
		return ErrKeyNotFound
	}

	if _, exists := s.data[newKey]; exists {
		return ErrKeyAlreadyExists
	}

	s.data[newKey] = value
	delete(s.data, oldKey)

	return nil
}

func (s *Store) Pop(key string) (string, error) {
	value, exists := s.data[key]

	if !exists {
		return "", ErrKeyNotFound
	}

	delete(s.data, key)

	return value, nil
}

func (s *Store) Len() int {
	return len(s.data)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}
