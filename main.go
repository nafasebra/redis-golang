package main

import (
	"fmt"
	"encoding/base64"
)

func main() {	
	data := NewStore()

	data.Set("a", "nafas")
	data.Set("b", "blah blah blah")

	data.Delete("b")

	fmt.Println(data.Get("a"))

	fmt.Println(data.Get("b"))
}

func SetKeyWithEncryption(store Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	return store.Get(key)
}
