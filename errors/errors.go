package errors

import "errors"

var (
	ErrEmptyKey         = errors.New("key is empty")
	ErrKeyNotFound      = errors.New("key not found")
	ErrKeyAlreadyExists = errors.New("key already exists")
	ErrInvalidKey       = errors.New("key cannot be empty")
	ErrInvalidExpiration = errors.New("The time is expired")
)
