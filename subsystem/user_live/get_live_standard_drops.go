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

// liveScore is nil if this is a skip
func getLiveStandardDrops(session *userdata.Session, liveScore *client.LiveScore, liveDifficulty *gamedata.LiveDifficulty) (generic.Array[client.LiveDropContent], bool) {
	drops := generic.Array[client.LiveDropContent]{}
	isRewardAccessoryInPresentBox := false
	// this is taken from https://suyo.be/sifas/wiki/gameplay/live-rewards
	// this is not in the database afaik
	dropCount := 10
	if liveDifficulty.LiveDifficultyType == enum.LiveDifficultyTypeHard {
		dropCount = 7
	} else if liveDifficulty.LiveDifficultyType == enum.LiveDifficultyTypeNormal {
		dropCount = 4
	}

	voltage := liveDifficulty.EvaluationSScore
	if liveScore != nil {
		voltage = liveScore.CurrentScore
	}

	if voltage >= liveDifficulty.EvaluationCScore {
		dropCount++
	}
	if voltage >= liveDifficulty.EvaluationBScore {
		dropCount++
	}
	if voltage >= liveDifficulty.EvaluationAScore {
		dropCount++
	}
	if voltage >= liveDifficulty.EvaluationSScore {
		dropCount++
	}

	if liveDifficulty.UnlockPattern == enum.LiveUnlockPatternStoryOnly {
		dropCount /= 2 // story songs give half the drop
	}
	for i := 0; i < dropCount; i++ {
		isRare := rand.Int31n(10000) < liveDifficulty.RareDropRate
		var content client.Content
		if isRare {
			content = liveDifficulty.RareDropContentGroup.GetRandomItemByDropColor(enum.NoteDropColorBronze)
		} else {
			content = liveDifficulty.DropContentGroup.GetRandomItemByDropColor(enum.NoteDropColorBronze)
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
