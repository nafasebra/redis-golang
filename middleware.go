package main

import (
	"log"
	"redis-golang/store"
	"time"
)

var _ store.Storer = (*LoggingMiddleware)(nil)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

// NewLoggingMiddleware wraps a store and logs its operations.
// A nil logger uses the standard logger.
func NewLoggingMiddleware(inner store.Storer, logger *log.Logger) *LoggingMiddleware {
	if logger == nil {
		logger = log.Default()
	}
	return &LoggingMiddleware{inner: inner, logger: logger}
}

func (m *LoggingMiddleware) Get(key string) (string, error) {
	start := time.Now()
	value, err := m.inner.Get(key)
	m.logger.Printf("operation=Get key=%q duration=%s error=%v", key, time.Since(start), err)
	return value, err
}

func (m *LoggingMiddleware) Set(key, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.logger.Printf("operation=Set key=%q duration=%s error=%v", key, time.Since(start), err)
	return err
}

func (m *LoggingMiddleware) Delete(key string) error {
	start := time.Now()
	err := m.inner.Delete(key)
	m.logger.Printf("operation=Delete key=%q duration=%s error=%v", key, time.Since(start), err)
	return err
}

func (m *LoggingMiddleware) Keys() []string {
	start := time.Now()
	keys := m.inner.Keys()
	m.logger.Printf("operation=Keys count=%d duration=%s", len(keys), time.Since(start))
	return keys
}

func (m *LoggingMiddleware) Len() int {
	start := time.Now()
	count := m.inner.Len()
	m.logger.Printf("operation=Len count=%d duration=%s", count, time.Since(start))
	return count
}
