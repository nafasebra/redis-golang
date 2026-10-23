package store_test

import (
	"errors"
	storeerrors "redis-golang/errors"
	"redis-golang/store"
	"redis-golang/store/kv"
	"redis-golang/store/ttl"
	"testing"
)

func TestStorer(t *testing.T) {
	for name, data := range map[string]store.Storer{"kv": kv.NewStore(), "ttl": ttl.NewTTLStore()} {
		t.Run(name, func(t *testing.T) {
			if err := data.Set("name", "Nafas"); err != nil {
				t.Fatal(err)
			}
			if value, err := data.Get("name"); err != nil || value != "Nafas" {
				t.Fatalf("Get = %q, %v", value, err)
			}
			if data.Len() != 1 || len(data.Keys()) != 1 {
				t.Fatal("expected one key")
			}
			if err := data.Delete("name"); err != nil {
				t.Fatal(err)
			}
			if _, err := data.Get("name"); !errors.Is(err, storeerrors.ErrKeyNotFound) {
				t.Fatalf("expected missing key error, got %v", err)
			}
		})
	}
}
