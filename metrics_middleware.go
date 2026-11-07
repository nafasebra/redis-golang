package main

import (
	"log"
	"redis-golang/store"
)

var _ store.Storer = (*MetricsMiddleware)(nil)

type MetricsMiddleware struct {
	inner store.Storer

	// count of creation (set)
	createCount int
	// count of selection (get)
	selectCount int
	// count of deletion (delete)
	deleteCount int
	// count of failed store operations
	failureCount int
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	m.selectCount++
	log.Printf("metrics operation=Get select_count=%d", m.selectCount)
	value, err := m.inner.Get(key)
	m.recordFailure("Get", err)
	return value, err
}

func (m *MetricsMiddleware) Set(key, value string) error {
	m.createCount++
	log.Printf("metrics operation=Set create_count=%d", m.createCount)
	err := m.inner.Set(key, value)
	m.recordFailure("Set", err)
	return err
}

func (m *MetricsMiddleware) Delete(key string) error {
	m.deleteCount++
	log.Printf("metrics operation=Delete delete_count=%d", m.deleteCount)
	err := m.inner.Delete(key)
	m.recordFailure("Delete", err)
	return err
}

func (m *MetricsMiddleware) recordFailure(operation string, err error) {
	if err == nil {
		return
	}

	m.failureCount++
	log.Printf("metrics operation=%s failure_count=%d error=%v", operation, m.failureCount, err)
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Len() int {
	return m.inner.Len()
}
