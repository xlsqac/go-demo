package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type memCache struct {
	maxMemSize    int64                     // 最大内存大小
	maxMemSizeStr string                    // 最大内存大小的 string 表示
	currMemSize   int64                     // 当前内存大小
	values        map[string]*memCacheValue // 缓存值
	locker        sync.RWMutex              // 锁
}

type memCacheValue struct {
	val        any           // 值
	expireTime time.Time     // 到期时间
	expire     time.Duration // 有效时长
	size       int64         // 占用内存大小
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemSize, mc.maxMemSizeStr = ParseSize(size)
	fmt.Println(mc.maxMemSize, mc.maxMemSizeStr)
	return true
}

// Set 设置 Key 的值
func (mc *memCache) Set(key string, val any, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValueSize(val),
	}
	if _, ok := mc.get(key); ok {
		mc.del(key)
	}
	mc.add(key, v)

	if mc.currMemSize > mc.maxMemSize {
		mc.del(key)
		log.Println(fmt.Sprintf("max memory size %s", mc.maxMemSizeStr))
		return false
	}

	return true
}

// Get 获取 Key 对应的值
func (mc *memCache) Get(key string) (any, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	val, ok := mc.get(key)
	if ok {
		if val.expire != 0 && val.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}
		return val.val, ok
	}
	return nil, false
}

// Del 删除 Key
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

// Exists 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.get(key)
	return ok
}

// Flush 清空缓存，删除所有 key
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemSize = 0
	return true
}

// Keys 获取 key 的数量
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) del(key string) {
	val, ok := mc.get(key)
	if ok && val != nil {
		mc.currMemSize -= val.size
		delete(mc.values, key)
	}

}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemSize += val.size
}

// NewMemCache 初始化 memCache
func NewMemCache() Cache {
	// 保存的值是一个 map，所以在初始化结构体的时候需要先初始化这个 map
	mc := &memCache{
		values: make(map[string]*memCacheValue),
	}
	return mc
}

func GetValueSize(val any) int64 {
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}
