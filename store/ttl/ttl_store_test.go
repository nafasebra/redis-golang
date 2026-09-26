package ttl

import (
	"testing"
	"time"

	"redis-golang/errors"
)

func TestNewTTLStore(t *testing.T) {
	store := NewTTLStore()

	if store == nil {
		t.Fatal("expected store to be created")
	}

	if store.data == nil {
		t.Fatal("expected data map to be initialized")
	}

	if store.Len() != 0 {
		t.Fatalf("expected empty store, got %d", store.Len())
	}
}

func TestSet(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("name", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	entry, exists := store.data["name"]

	if !exists {
		t.Fatal("expected key to exist")
	}

	if entry.value != "Nafas" {
		t.Fatalf("expected value Nafas, got %s", entry.value)
	}

	if !entry.expire_at.Equal(expireAt) {
		t.Fatal("expected expiration time to be saved correctly")
	}
}

func TestSetEmptyKey(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("", "value", expireAt)

	if err != errors.ErrEmptyKey {
		t.Fatalf("expected ErrEmptyKey, got %v", err)
	}
}

func TestSetInvalidExpiration(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(-time.Minute)

	err := store.Set("name", "Nafas", expireAt)

	if err != errors.ErrInvalidExpiration {
		t.Fatalf("expected ErrInvalidExpiration, got %v", err)
	}
}

func TestSetOverwrite(t *testing.T) {
	store := NewTTLStore()

	firstExpiration := time.Now().Add(time.Minute)
	secondExpiration := time.Now().Add(2 * time.Minute)

	err := store.Set("name", "Nafas", firstExpiration)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.Set("name", "Ali", secondExpiration)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	value, err := store.GetWithTTL("name")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != "Ali" {
		t.Fatalf("expected Ali, got %s", value)
	}
}

func TestGetWithTTL(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("name", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	value, err := store.GetWithTTL("name")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if value != "Nafas" {
		t.Fatalf("expected Nafas, got %s", value)
	}
}

func TestGetWithTTLEmptyKey(t *testing.T) {
	store := NewTTLStore()

	_, err := store.GetWithTTL("")

	if err != errors.ErrEmptyKey {
		t.Fatalf("expected ErrEmptyKey, got %v", err)
	}
}

func TestGetWithTTLNonExistingKey(t *testing.T) {
	store := NewTTLStore()

	_, err := store.GetWithTTL("name")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}
}

func TestGetWithTTLExpiredKey(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(50 * time.Millisecond)

	err := store.Set("name", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	_, err = store.GetWithTTL("name")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}

	if _, exists := store.data["name"]; exists {
		t.Fatal("expected expired key to be deleted")
	}
}

func TestDelete(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("name", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.Delete("name")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, exists := store.data["name"]; exists {
		t.Fatal("expected key to be deleted")
	}
}

func TestDeleteNonExistingKey(t *testing.T) {
	store := NewTTLStore()

	err := store.Delete("name")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}
}

func TestRename(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("old", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.Rename("old", "new")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, oldExists := store.data["old"]

	if oldExists {
		t.Fatal("expected old key to be deleted")
	}

	entry, newExists := store.data["new"]

	if !newExists {
		t.Fatal("expected new key to exist")
	}

	if entry.value != "Nafas" {
		t.Fatalf("expected Nafas, got %s", entry.value)
	}
}

func TestRenameEmptyOldKey(t *testing.T) {
	store := NewTTLStore()

	err := store.Rename("", "new")

	if err != errors.ErrInvalidKey {
		t.Fatalf("expected ErrInvalidKey, got %v", err)
	}
}

func TestRenameEmptyNewKey(t *testing.T) {
	store := NewTTLStore()

	err := store.Rename("old", "")

	if err != errors.ErrInvalidKey {
		t.Fatalf("expected ErrInvalidKey, got %v", err)
	}
}

func TestRenameNonExistingKey(t *testing.T) {
	store := NewTTLStore()

	err := store.Rename("old", "new")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}
}

func TestRenameExpiredKey(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(50 * time.Millisecond)

	err := store.Set("old", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	err = store.Rename("old", "new")

	if err != errors.ErrKeyNotFound {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}

	if _, exists := store.data["old"]; exists {
		t.Fatal("expected expired old key to be deleted")
	}
}

func TestRenameToExistingKey(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	err := store.Set("old", "Nafas", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.Set("new", "Ali", expireAt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.Rename("old", "new")

	if err != errors.ErrKeyAlreadyExists {
		t.Fatalf("expected ErrKeyAlreadyExists, got %v", err)
	}
}

func TestLen(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	_ = store.Set("name", "Nafas", expireAt)
	_ = store.Set("age", "25", expireAt)

	if store.Len() != 2 {
		t.Fatalf("expected length 2, got %d", store.Len())
	}
}

func TestLenAfterDelete(t *testing.T) {
	store := NewTTLStore()

	expireAt := time.Now().Add(time.Minute)

	_ = store.Set("name", "Nafas", expireAt)
	_ = store.Set("age", "25", expireAt)

	_ = store.Delete("name")

	if store.Len() != 1 {
		t.Fatalf("expected length 1, got %d", store.Len())
	}
}
