package core

import (
	"sync"
	"time"
)

type Idempotency struct {
	mu    sync.Mutex
	locks map[string]time.Time
	ttl   time.Duration
}

func NewIdempotency(ttl time.Duration) *Idempotency {
	return &Idempotency{
		locks: make(map[string]time.Time),
		ttl:   ttl,
	}
}

func (i *Idempotency) TryLock(key string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	if t, ok := i.locks[key]; ok && time.Since(t) < i.ttl {
		return false
	}
	i.locks[key] = time.Now()
	return true
}

func (i *Idempotency) Unlock(key string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.locks, key)
}
