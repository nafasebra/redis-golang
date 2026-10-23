package main

import (
	"encoding/base64"
	"fmt"
	"redis-golang/store"
	"redis-golang/store/ttl"
	"time"
)

func main() {
	data := ttl.NewTTLStore()

	if err := data.Set("name", "Nafas"); err != nil {
		fmt.Println(err)
		return
	}
	if err := data.SetWithTTL("code", "1234", time.Now().Add(time.Minute)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("code"))
	fmt.Println("keys:", data.Keys())
}

func SetKeyWithEncryption(data store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := data.Set(key, encoded); err != nil {
		return "", err
	}
	return data.Get(key)
}
