package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// -----------------------------
// RedisCache implement Cache interface
// -----------------------------
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache constructor
func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		DialTimeout:  2 * time.Second,
		PoolTimeout:  3 * time.Second,
	})
	return &RedisCache{client: rdb}
}

// Set value with TTL
func (r *RedisCache) Set(ctx context.Context, key string, value []byte, ttlSec int) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return r.client.Set(ctx, key, value, time.Duration(ttlSec)*time.Second).Err()
}

// Get value
func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	val, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil // ไม่ error แต่ key ไม่มี
	}
	return val, err
}

// Delete key
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return r.client.Del(ctx, key).Err()
}

// Keys ดึง key ตาม pattern
func (r *RedisCache) Keys(ctx context.Context, pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return r.client.Keys(ctx, pattern).Result()
}

// Ping health check
func (r *RedisCache) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	_, err := r.client.Ping(ctx).Result()
	return err
}
