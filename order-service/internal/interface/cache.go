package port

import "context"

// -----------------------------
// Cache interface สำหรับ abstraction
// -----------------------------
type Cache interface {
	Set(ctx context.Context, key string, value []byte, ttlSec int) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	Ping(ctx context.Context) error
}
