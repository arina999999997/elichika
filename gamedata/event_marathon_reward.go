package gamedata

import (
	"elichika/client"
	"elichika/dictionary"

	"xorm.io/xorm"
)

func loadEventMarathonReward(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.EventMarathonReward = map[int32][]*client.Content{}
	for _, eventMarathon := range gamedata.EventMarathon {
		for i := range eventMarathon.TopStatus.EventMarathonRewardMasterRows.Slice {
			reward := &eventMarathon.TopStatus.EventMarathonRewardMasterRows.Slice[i]
			gamedata.EventMarathonReward[reward.RewardGroupId] = append(gamedata.EventMarathonReward[reward.RewardGroupId], &reward.RewardContent)
		}
	}
}

func init() {
	addLoadFunc(loadEventMarathonReward)
	addPrequisite(loadEventMarathonReward, loadEventMarathon)
}
