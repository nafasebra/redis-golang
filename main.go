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
	data := make(map[string]string)

	data["a"] = "nafas"
	data["b"] = "1234"

	value, exists := data["b"];

	fmt.Println(data["a"])
	fmt.Println(value, exists)


    fmt.Println("Hello, World!")
}
