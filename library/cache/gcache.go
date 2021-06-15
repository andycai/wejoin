package cache

import "github.com/bluele/gcache"

type Cache struct {
	cache gcache.Cache
}

func New(size int) *Cache {
	c := new(Cache)
	c.cache = gcache.New(size).
		LRU().
		Build()
	return c
}

func (c *Cache) Set(key, value interface{}) error {
	return c.cache.Set(key, value)
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	return c.cache.Get(key)
}

func (c *Cache) Has(key interface{}) bool {
	return c.cache.Has(key)
}

func (c *Cache) Keys(checkExpired bool) []interface{} {
	return c.cache.Keys(checkExpired)
}
