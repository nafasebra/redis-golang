package main

import "sort"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	value, exist := s.data[key];
	return value, exist
}

func (s *Store) Set(key, value string) {
	s.data[key] = value;
}

func (s *Store) Delete(key string) {
	delete(s.data, key);
}

func (s *Store) Keys() ([]string) {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings((keys))

	return keys
}

func (s *Store) Rename(oldkey, newkey string) {
	value := s.data[oldkey]
	s.data[newkey] = value
	delete(s.data, oldkey)
}

func (s *Store) Pop(key string) (string, bool) {
	value, exist := s.data[key]
	delete(s.data, key)
	return value, exist
}

func (s *Store) Len() int {
	return len(s.data)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

