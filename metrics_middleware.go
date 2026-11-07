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
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	m.selectCount++
	log.Printf("metrics operation=Get select_count=%d", m.selectCount)
	return m.inner.Get(key)
}

func (m *MetricsMiddleware) Set(key string, value string) error {
	m.createCount++
	log.Printf("metrics operation=Set create_count=%d", m.createCount)
	return m.inner.Set(key, value)
}

func (m *MetricsMiddleware) Delete(key string) error {
	m.deleteCount++
	log.Printf("metrics operation=Delete delete_count=%d", m.deleteCount)
	return m.inner.Delete(key)
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Len() int {
	return m.inner.Len()
}
