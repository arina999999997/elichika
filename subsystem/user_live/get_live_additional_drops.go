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

	dropContentGroup := liveDifficulty.AdditionalDropContentGroup
	rareDropContentGroup := liveDifficulty.AdditionalRareDropContentGroup
	if liveDifficulty.UnlockPattern == enum.LiveUnlockPatternDaily {
		// daily doesn't use additional group, as using them seems to result in wrong drop
		dropContentGroup, rareDropContentGroup = getDropContentGroups(session, liveScore, liveDifficulty)
	}

	for i := int32(0); i < dropCount; i++ {
		isRare := rand.Int31n(10000) < liveDifficulty.RareDropRate
		var content client.Content
		if isRare {
			content = rareDropContentGroup.GetRandomItemByDropColor(enum.NoteDropColorBronze)
		} else {
			content = dropContentGroup.GetRandomItemByDropColor(enum.NoteDropColorBronze)
		}

		result := user_content.AddContent(session, content)
		if content.ContentType == enum.ContentTypeAccessory {
			isRewardAccessoryInPresentBox = isRewardAccessoryInPresentBox || result.(bool)
		}
		drops.Append(client.LiveDropContent{
			DropColor: enum.NoteDropColorBronze,
			Content:   content,
			IsRare:    isRare,
		})
	}
	return drops, isRewardAccessoryInPresentBox
}
