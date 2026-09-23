package main

import "fmt"

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

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {	
	data := NewStore()

	data.Set("a", "nafas")
	data.Set("b", "blah blah blah")

	fmt.Println(data.Get("a"))
}
