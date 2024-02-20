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

func getLiveAdditionalDrops(session *userdata.Session, liveScore *client.LiveScore, liveDifficulty *gamedata.LiveDifficulty) (generic.Array[client.LiveDropContent], bool) {
	drops := generic.Array[client.LiveDropContent]{}

	isRewardAccessoryInPresentBox := false

	totalTechnique := int32(0)
	for _, liveCardStat := range liveScore.CardStatDict.Map {
		totalTechnique += liveCardStat.BaseParameter.Technique
	}

	dropCount := int32(0)
	for (totalTechnique > liveDifficulty.BottomTechnique) && (dropCount < liveDifficulty.AdditionalDropMaxCount) {
		totalTechnique -= liveDifficulty.AdditionalDropDecayTechnique
		dropCount++
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
