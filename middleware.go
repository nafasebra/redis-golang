package main

import (
	"log"
	"redis-golang/store"
	"time"
)

const (
	colorReset = "\033[0m"
	colorWhite = "\033[37m"
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
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
	m.logger.Printf(colorForError(err)+"operation=Get key=%q duration=%s error=%v"+colorReset, key, time.Since(start), err)
	return value, err
}

func (m *LoggingMiddleware) Set(key, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.logger.Printf(colorForError(err)+"operation=Set key=%q duration=%s error=%v"+colorReset, key, time.Since(start), err)
	return err
}

func (m *LoggingMiddleware) Delete(key string) error {
	start := time.Now()
	err := m.inner.Delete(key)
	m.logger.Printf(colorForError(err)+"operation=Delete key=%q duration=%s error=%v"+colorReset, key, time.Since(start), err)
	return err
}

func (m *LoggingMiddleware) Keys() []string {
	start := time.Now()
	keys := m.inner.Keys()
	m.logger.Printf(colorWhite+"operation=Keys count=%d duration=%s"+colorReset, len(keys), time.Since(start))
	return keys
}

func (m *LoggingMiddleware) Len() int {
	start := time.Now()
	count := m.inner.Len()
	m.logger.Printf(colorWhite+"operation=Len count=%d duration=%s"+colorReset, count, time.Since(start))
	return count
}

func colorForError(err error) string {
	if err != nil {
		return colorRed
	}
	return colorWhite
}
