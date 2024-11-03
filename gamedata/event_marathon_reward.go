package gamedata

import (
	"elichika/client"
)

func loadEventMarathonReward(gamedata *Gamedata) {
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
