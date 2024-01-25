package cache

// this package implement a generic cache handler that is thread safe
// the behaviour is as follow:
// - an expiration timestamp is set for the caching object
// - whenever a request come in, it will check if the cache is valid, if it is then it's returned,
// otherwise it will trigger a recalculation
// - if a recalculation is being done then no more calculation will start
//
// - basically the thread should acquire the lock, then make a copy of the pointer to data, then release the lock if the data is valid
// - otherwise it should recalculate then relrease the lock
import (
	"sync"
)

type CachedObject[T any] struct {
	Value    *T
	ExpireAt int64
	mutex    sync.Mutex
}

func (co *CachedObject[T]) Acquire() {
	co.mutex.Lock()
}

func (co *CachedObject[T]) Release() {
	co.mutex.Unlock()
}
