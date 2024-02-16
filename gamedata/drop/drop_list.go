package drop

import (
	"elichika/client"

	"math/rand"
)

// equal weighted drop list
type DropList struct {
	contents []client.Content
	n        int32
}

func (dl *DropList) AddContent(content client.Content) {
	dl.contents = append(dl.contents, content)
	dl.n++
}

func (dl *DropList) GetRandomDrop() client.Content {
	return dl.contents[rand.Int31n(dl.n)]
}

func (dl *DropList) GetRandomDrops(count int32) []client.Content {
	result := []client.Content{}
	for i := int32(0); i < count; i++ {
		result = append(result, dl.GetRandomDrop())
	}
	return result
}
