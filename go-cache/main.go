// package main
package main

import "time"

// Cache
type Cache struct {
	Data map[string]any // 存放数据，支持各种数据
	// 和 redis 惰性删除一样，key 过期时不会立即删除，而是在下次访问时删除
	ttl map[string]int // 存放过期时间
}

// exist
// 1. 检查 key 是否存在 0 不存在 1 存在
func (c *Cache) exist(key string) int {

}

// get
// 1. 检查 key 是否存在
// 2. 检查 key 是否过期
func (c *Cache) get(key string) any {
	v, ok := c.Data[key]
	if ok {
		if !c.checkTTL(key) {
			return nil
		}
		return v
	}
	return nil
}

func (c *Cache) set(key string, value any, ex int) {}

func (c *Cache) del(key string) {}

func (c *Cache) checkTTL(key string) bool {
	ttl, ok := c.ttl[key]
	if !ok {
		return true
	} else {
		if ttl < c.now() {
			return false
		}
	}
	return true
}

func (c *Cache) now() int {
	currentTime := int(time.Now().Unix())
	return currentTime
}

// main
func main() {

}
