package gamedata

import (
	"elichika/client"
)

func loadEventMiningReward(gamedata *Gamedata) {
	gamedata.EventMiningReward = map[int32][]*client.Content{}
	for _, eventMining := range gamedata.EventMining {
		for i := range eventMining.TopStatus.EventMiningRewardMasterRows.Slice {
			reward := &eventMining.TopStatus.EventMiningRewardMasterRows.Slice[i]
			gamedata.EventMiningReward[reward.RewardGroupId] = append(gamedata.EventMiningReward[reward.RewardGroupId], &reward.RewardContent)
		}
	}
}

func init() {
	addLoadFunc(loadEventMiningReward)
	addPrequisite(loadEventMiningReward, loadEventMining)
}
