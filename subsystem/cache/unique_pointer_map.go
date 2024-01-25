package cache

import (
	"sync"
)

// the cached map is used to setup element for the map of pointer,
// making sure that no 2 pointers will be created for a single key

// This can be optimized by initialising all the possible keys, but let's not think too hard right now

type UniquePointerMap[K int32 | int64, V any] struct {
	m  map[K]*V
	mu sync.Mutex
}

func (ump *UniquePointerMap[K, V]) Get(key K) *V {
	ump.mu.Lock()
	defer ump.mu.Unlock()
	if ump.m == nil {
		ump.m = make(map[K]*V)
	}
	ptr, exist := ump.m[key]
	if !exist {
		ptr = new(V)
		ump.m[key] = ptr
	}
	return ptr
}
