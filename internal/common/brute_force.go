package common

import (
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
)

const MaxLoginAttempts = 5
const LoginBlockMinutes = 10

func RegisterFailedLogin(email string) {
	key := fmt.Sprintf("failed_login:%s", email)

	// Attempt increment
	_, err := Cache.Increment(key, 1)

	if err == memcache.ErrCacheMiss {
		// Key does not exist → create it with value=1 and TTL
		Cache.Set(&memcache.Item{
			Key:        key,
			Value:      []byte("1"),
			Expiration: int32(LoginBlockMinutes * 60),
		})
		return
	}

	// If memcached fails, ignore (fail open)
	if err != nil {
		return
	}
}

func IsBlocked(email string) bool {
	key := fmt.Sprintf("failed_login:%s", email)

	item, err := Cache.Get(key)
	if err != nil {
		return false // no key = no failures
	}

	// Parse stored integer
	count, _ := strconv.Atoi(string(item.Value))

	return count >= MaxLoginAttempts
}

func ResetLoginFailures(email string) {
	key := fmt.Sprintf("failed_login:%s", email)
	Cache.Delete(key)
}
