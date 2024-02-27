package user_live

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

// how gimmick (note) drop system is implemented:
// - when starting live, DropChooseCount notes are randomly chosen and given a drop color.
// - if these note is hit, the user is entitled to a drop
// - this might not be the totally correct interpretation

func getLiveGimmickDrops(session *userdata.Session, liveStage *client.LiveStage, liveScore *client.LiveScore, liveDifficulty *gamedata.LiveDifficulty) (generic.Array[client.LiveDropContent], bool) {
	drops := generic.Array[client.LiveDropContent]{}

	isRewardAccessoryInPresentBox := false

	for i := range liveStage.LiveNotes.Slice {
		if liveStage.LiveNotes.Slice[i].NoteRandomDropColor != enum.NoteDropColorNon {
			id := liveStage.LiveNotes.Slice[i].Id
			// this note has a drop, see if the user actually cleared it
			if liveScore.ResultDict.Map[id].JudgeType >= enum.JudgeTypeBad {
				// add a drop
				content := liveDifficulty.NoteDropGroup.GetRandomItemByDropColor(liveStage.LiveNotes.Slice[i].NoteRandomDropColor)
				result := user_content.AddContent(session, content)
				if content.ContentType == enum.ContentTypeAccessory {
					isRewardAccessoryInPresentBox = isRewardAccessoryInPresentBox || result.(bool)
				}
				drops.Append(client.LiveDropContent{
					DropColor: liveStage.LiveNotes.Slice[i].NoteRandomDropColor,
					Content:   content,
					IsRare:    false,
				})
			}
		}
	}
	return drops, isRewardAccessoryInPresentBox
}
