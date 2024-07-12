package memcache

import (
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemCache struct {
	client *memcache.Client
	url    string
	sync.RWMutex
}

func Connect(url string) *MemCache {
	return &MemCache{
		client: memcache.New(url),
		url:    url,
	}
}

func (cache *MemCache) Close() error {
	cache.Lock()
	defer cache.Unlock()
	if cache.client != nil {
		err := cache.client.Close()
		cache.client = nil
		return err
	}
	return nil
}

func (cache *MemCache) Get(key string) (*memcache.Item, error) {
	cache.Lock()
	defer cache.Unlock()
	if cache.client == nil {
		cache.reconnect()
	}
	item, err := cache.client.Get(key)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (cache *MemCache) Set(key string, value []byte) error {
	cache.Lock()
	defer cache.Unlock()
	if cache.client == nil {
		cache.reconnect()
	}
	cacheErr := cache.client.Set(&memcache.Item{Key: key, Value: value, Expiration: 60})
	if cacheErr != nil {
		return cacheErr
	}
	return nil
}
func (cache *MemCache) reconnect() error {
	cache.Close()
	cache.Lock()
	defer cache.Unlock()
	cache.client = memcache.New(cache.url)
	return nil
}
