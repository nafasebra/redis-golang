package main

import "testing"

func TestStore(t *testing.T) {
	store := NewStore()

	// Test Set
	store.Set("a", "nafas")

	// Test Get
	value, exists := store.Get("a")

	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "nafas" {
		t.Fatalf("expected nafas, got %s", value)
	}

	// Test Delete
	store.Delete("a")

	_, exists = store.Get("a")

	if exists {
		t.Fatal("expected key to be deleted")
	}
}
