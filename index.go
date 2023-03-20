package cache_memory

import (
	"github.com/infrago/cache"
)

func Driver() cache.Driver {
	return &memoryDriver{}
}

func init() {
	cache.Register("memory", Driver())
}
