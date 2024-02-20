package user_live

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/userdata"

	"math/rand"
)

func getSkipAdditionalDrops(session *userdata.Session, roundedUp bool, liveDifficulty *gamedata.LiveDifficulty) (generic.Array[client.LiveDropContent], bool) {
	drops := generic.Array[client.LiveDropContent]{}

	isRewardAccessoryInPresentBox := false

	dropCount := liveDifficulty.AdditionalDropMaxCount / 2
	if roundedUp {
		dropCount = (liveDifficulty.AdditionalDropMaxCount + 1) / 2
	}

	for i := int32(0); i < dropCount; i++ {
		isRare := rand.Int31n(10000) < liveDifficulty.RareDropRate
		var content client.Content
		if isRare {
			content = liveDifficulty.AdditionalRareDropContentGroup.GetRandomItem()
		} else {
			content = liveDifficulty.AdditionalDropContentGroup.GetRandomItem()
		}

		result := user_content.AddContent(session, content)
		if content.ContentType == enum.ContentTypeAccessory {
			isRewardAccessoryInPresentBox = isRewardAccessoryInPresentBox || result.(bool)
		}
		drops.Append(client.LiveDropContent{
			DropColor: enum.NoteDropColorBronze, // not sure if this still do anything
			Content:   content,
			IsRare:    isRare,
		})
	}
	return drops, isRewardAccessoryInPresentBox
}
