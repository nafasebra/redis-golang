package main

import (
	"encoding/base64"
	"fmt"
	"redis-golang/store"
	"redis-golang/store/kv"
)

func main() {
	sdata := kv.NewStore()
	metrics := NewMetricsMiddleware(sdata)
	logger := NewLoggingMiddleware(metrics, nil)

	if err := logger.Set("name", "Nafas"); err != nil {
		fmt.Println(err)
		return
	}
	if err := logger.Set("code", "1234"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(logger.Get("name"))
	fmt.Println(logger.Get("code"))
	fmt.Println("keys:", logger.Keys())

	if err := logger.Delete("code"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("keys after delete:", logger.Keys())
}

func SetKeyWithEncryption(data store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := data.Set(key, encoded); err != nil {
		return "", err
	}
	return data.Get(key)
}
