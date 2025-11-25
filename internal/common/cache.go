package common

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var Cache *memcache.Client

func InitCache() {
	// Memcached default port is 11211
	Cache = memcache.New("memcached:11211")

	err := Cache.Ping()
	if err != nil {
		log.Fatal("Memcached connection failed:", err)
	}

	log.Println("Memcached cache connected")
}

func CacheGet(key string, dest interface{}) bool {
	item, err := Cache.Get(key)
	if err != nil {
		return false
	}
	return json.Unmarshal(item.Value, dest) == nil
}

func CacheSet(key string, val interface{}, ttl time.Duration) {
	data, _ := json.Marshal(val)

	Cache.Set(&memcache.Item{
		Key:        key,
		Value:      data,
		Expiration: int32(ttl.Seconds()),
	})
}
