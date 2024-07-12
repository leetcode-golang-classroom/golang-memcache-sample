package types

import "github.com/bradfitz/gomemcache/memcache"

type MemCache interface {
	Close() error
	Get(key string) (*memcache.Item, error)
	Set(key string, data []byte) error
}
