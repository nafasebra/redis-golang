package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string)
	}
}

func main() {
    fmt.Println("Hello, World!")
}
