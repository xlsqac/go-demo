// package main
package main

import "xlsqac/gocache/cache"

// main
func main() {
	c := cache.NewMemCache()
	c.SetMaxMemory("100MB")
	c.Set("int", 1, 0)
	c.Set("bool", false, 0)
	c.Set("data", map[string]interface{}{"a": 1}, 0)
	c.Get("int")
	c.Del("int")
	c.Flush()
	c.Keys()
	cache.ParseSize("100MB")
}
