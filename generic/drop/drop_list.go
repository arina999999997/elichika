package drop

import (
	"math/rand"
)

// equal weighted drop list
type DropList[T any] struct {
	items []T
	n     int32
}

func (dl *DropList[T]) AddItem(t T) {
	dl.items = append(dl.items, t)
	dl.n++
}

func (dl *DropList[T]) GetRandomItem() T {
	return dl.items[rand.Int31n(dl.n)]
}

func (dl *DropList[T]) GetRandomItems(count int32) []T {
	result := []T{}
	for i := int32(0); i < count; i++ {
		result = append(result, dl.GetRandomItem())
	}
	return result
}
