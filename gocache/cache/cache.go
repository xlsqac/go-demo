package cache

import "time"

type Cache interface {
	SetMaxMemory(size string) bool
	Set(key string, val any, expire time.Duration) bool
	Get(key string) (any, bool)
	Del(key string) bool
	Exists(key string) bool
	Flush() bool
	Keys() int64
}
