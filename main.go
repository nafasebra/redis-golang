package main

import (
	"encoding/base64"
	"fmt"
	"redis-golang/store"
	"redis-golang/store/kv"
	"redis-golang/store/ttl"
)

func main() {
	stores := []store.Storer{kv.NewStore(), ttl.NewTTLStore()}
	
	for _, data := range stores {
		value, err := SetKeyWithEncryption(data, "a", "nafas")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(value)
	}
}

func SetKeyWithEncryption(data store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := data.Set(key, encoded); err != nil {
		return "", err
	}
	return data.Get(key)
}
