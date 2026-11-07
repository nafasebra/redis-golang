package main

import (
	"log"
	"redis-golang/store"
	"time"
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
	// total time spent handling store operations
	totalLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	m.selectCount++
	log.Printf(colorGreen+"metrics operation=Get select_count=%d"+colorReset, m.selectCount)
	start := time.Now()
	value, err := m.inner.Get(key)
	m.recordLatency("Get", time.Since(start))
	m.recordFailure("Get", err)
	return value, err
}

func (m *MetricsMiddleware) Set(key, value string) error {
	m.createCount++
	log.Printf(colorGreen+"metrics operation=Set create_count=%d"+colorReset, m.createCount)
	start := time.Now()
	err := m.inner.Set(key, value)
	m.recordLatency("Set", time.Since(start))
	m.recordFailure("Set", err)
	return err
}

func (m *MetricsMiddleware) Delete(key string) error {
	m.deleteCount++
	log.Printf(colorGreen+"metrics operation=Delete delete_count=%d"+colorReset, m.deleteCount)
	start := time.Now()
	err := m.inner.Delete(key)
	m.recordLatency("Delete", time.Since(start))
	m.recordFailure("Delete", err)
	return err
}

func (m *MetricsMiddleware) recordLatency(operation string, latency time.Duration) {
	m.totalLatency += latency
	log.Printf(colorGreen+"metrics operation=%s latency=%s total_latency=%s"+colorReset, operation, latency, m.totalLatency)
}

func (m *MetricsMiddleware) recordFailure(operation string, err error) {
	if err == nil {
		return
	}

	m.failureCount++
	log.Printf(colorRed+"metrics operation=%s failure_count=%d error=%v"+colorReset, operation, m.failureCount, err)
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Len() int {
	return m.inner.Len()
}
