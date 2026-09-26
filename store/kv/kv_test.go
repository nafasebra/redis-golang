package kv

import (
	"testing"

	"redis-golang/errors"
)

func TestSet(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "Nafas")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	value, exists := store.Get("name")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "Nafas" {
		t.Fatalf("expected Nafas, got %s", value)
	}
}

func TestSetInvalidKey(t *testing.T) {
	store := NewStore()

	err := store.Set("", "value")

	if err != errors.ErrInvalidKey {
		t.Fatalf("expected ErrInvalidKey, got %v", err)
	}
}

func TestSetOverwrite(t *testing.T) {
	store := NewStore()

	_ = store.Set("name", "Nafas")
	_ = store.Set("name", "Ali")

	value, exists := store.Get("name")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "Ali" {
		t.Fatalf("expected Ali, got %s", value)
	}
}

func TestGetExistingKey(t *testing.T) {
	store := NewStore()

	_ = store.Set("name", "Nafas")

	value, exists := store.Get("name")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "Nafas" {
		t.Fatalf("expected Nafas, got %s", value)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	store := NewStore()

	value, exists := store.Get("name")

	if exists {
		t.Fatal("expected key not to exist")
	}

	if value != "" {
		t.Fatalf("expected empty value, got %s", value)
	}
}

func TestDelete(t *testing.T) {
	store := NewStore()

	_ = store.Set("name", "Nafas")

	err := store.Delete("name")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, exists := store.Get("name")

	if exists {
		t.Fatal("expected key to be deleted")
	}
}

func TestDeleteNonExistingKey(t *testing.T) {
	store := NewStore()

	err := store.Delete("name")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}
}

func TestKeys(t *testing.T) {
	store := NewStore()

	_ = store.Set("banana", "1")
	_ = store.Set("apple", "2")
	_ = store.Set("orange", "3")

	keys := store.Keys()

	expected := []string{
		"apple",
		"banana",
		"orange",
	}

	if len(keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(keys))
	}

	for i := range expected {
		if keys[i] != expected[i] {
			t.Fatalf("expected key %s, got %s", expected[i], keys[i])
		}
	}
}

func TestKeysEmptyStore(t *testing.T) {
	store := NewStore()

	keys := store.Keys()

	if len(keys) != 0 {
		t.Fatalf("expected 0 keys, got %d", len(keys))
	}
}

func TestLen(t *testing.T) {
	store := NewStore()

	_ = store.Set("name", "Nafas")
	_ = store.Set("age", "25")

	if store.Len() != 2 {
		t.Fatalf("expected length 2, got %d", store.Len())
	}
}

func TestLenAfterDelete(t *testing.T) {
	store := NewStore()

	_ = store.Set("name", "Nafas")
	_ = store.Set("age", "25")

	_ = store.Delete("name")

	if store.Len() != 1 {
		t.Fatalf("expected length 1, got %d", store.Len())
	}
}
