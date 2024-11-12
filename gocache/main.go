// package main
package main

import (
	cache_server "xlsqac/gocache/cache-server"
)

// main
func main() {
	c := cache_server.NewMemCache()
	c.SetMaxMemory("100MB")
	c.Set("int", 1, 0)
	c.Set("bool", false, 0)
	c.Set("data", 1)
	c.Get("int")
	c.Del("int")
	c.Flush()
	c.Keys()
}
