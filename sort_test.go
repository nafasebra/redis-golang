package main

import "testing"

func TestSet(t *testing.T) {
	store := NewStore()

	store.Set("a", "nafas")

	value, exists := store.Get("a")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "nafas" {
		t.Fatalf("expected nafas, got %s", value)
	}
}

func TestSetOverwrite(t *testing.T) {
	store := NewStore()

	store.Set("a", "first")
	store.Set("a", "second")

	value, exists := store.Get("a")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "second" {
		t.Fatalf("expected second, got %s", value)
	}
}

func TestGetExistingKey(t *testing.T) {
	store := NewStore()

	store.Set("a", "nafas")

	value, exists := store.Get("a")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "nafas" {
		t.Fatalf("expected nafas, got %s", value)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	store := NewStore()

	value, exists := store.Get("a")

	if exists {
		t.Fatal("expected key to not exist")
	}

	if value != "" {
		t.Fatalf("expected empty value, got %s", value)
	}
}

func TestDeleteExistingKey(t *testing.T) {
	store := NewStore()

	store.Set("a", "nafas")
	store.Delete("a")

	_, exists := store.Get("a")

	if exists {
		t.Fatal("expected key to be deleted")
	}
}

func TestDeleteNonExistingKey(t *testing.T) {
	store := NewStore()

	store.Delete("a")
}

func TestKeys(t *testing.T) {
	store := NewStore()

	store.Set("c", "hello")
	store.Set("a", "nafas")
	store.Set("b", "blah")

	keys := store.Keys()

	expected := []string{"a", "b", "c"}

	if len(keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(keys))
	}

	for i := range expected {
		if keys[i] != expected[i] {
			t.Fatalf("expected %v, got %v", expected, keys)
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
