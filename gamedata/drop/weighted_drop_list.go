package drop

import (
	"elichika/client"

	"math/rand"
)

// equal weighted drop list
type WeightedDropList struct {
	contents    []client.Content
	weights     []int32
	totalWeight int32
	n           int
}

func (wdl *WeightedDropList) AddContent(content client.Content, weight int32) {
	wdl.contents = append(wdl.contents, content)
	wdl.totalWeight += weight
	wdl.weights = append(wdl.weights, wdl.totalWeight)
	wdl.n++
}

func (wdl *WeightedDropList) GetRandomDrop() client.Content {
	value := rand.Int31n(wdl.totalWeight)
	low := 0
	high := wdl.n - 1
	var mid, res int
	for low <= high {
		mid = (low + high) / 2
		if wdl.weights[mid] < value {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return wdl.contents[res]
}

func (wdl *WeightedDropList) GetRandomDrops(count int32) []client.Content {
	result := []client.Content{}
	for i := int32(0); i < count; i++ {
		result = append(result, wdl.GetRandomDrop())
	}
	return result
}
