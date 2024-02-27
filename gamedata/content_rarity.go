// define content rarity to decide the drop rate from live.
// Generate it using the follow method:
// - Iterate through a list of defined cost tables
// - Add up all the cost, to represent how much we would need to "max out" an account
// - The drop rate of an item is just proportional to the amount we need
// - Finally override the items with a list from "s_content_rarity" in serverdata.db if it is there TODO: actually do this
// - Note that this ignore the relative drop from training, for example (because it's fixed), but it should works well enough
//
// This data can then be used as the weight of each possible drop
// - the core idea is that for a drop list, each time we roll it, we should get the same amount of "progression"
// - if we need "x" of an item in total, and we get "a" from a drop, then the progression is defined to be "a" / "x"
// - so the relative chance to get an item should balance the progression of the drop
// - picking the weight of each type to be "x" / "a" should works
// - the expected progression gain per drop is then always ("x" / "a") / ("sum of weight") * "a" / "x" = 1 / "sum of weight"
// - In practice, we have some adjustment to make things feel better

package gamedata

import (
	"elichika/dictionary"
	"elichika/enum"
	"elichika/item"

	"fmt"

	"xorm.io/xorm"
)

type ContentRarity struct {
	RarityTable map[int32]map[int32]int64
}

func (cr *ContentRarity) AddContent(contentType int32, contentId int32, contentAmount int32) {
	if contentType == enum.ContentTypeGameMoney || contentType == enum.ContentTypeCardExp {
		// there are other ways of obtaining these, we need to adjust the chance or it's very hard to get other things
		contentAmount /= 20
		if contentType == enum.ContentTypeGameMoney {
			contentId = item.Gold.ContentId
		}
	}
	_, exist := cr.RarityTable[contentType]
	if !exist {
		cr.RarityTable[contentType] = map[int32]int64{}
	}
	_, exist = cr.RarityTable[contentType][contentId]
	if !exist {
		cr.RarityTable[contentType][contentId] = int64(contentAmount)
	} else {
		cr.RarityTable[contentType][contentId] += int64(contentAmount)
	}
}

func (cr *ContentRarity) GetWeight(contentType int32, contentId int32, contentAmount int32) int32 {
	if contentType == enum.ContentTypeExchangeEventPoint { // this is no longer used
		return 0
	}
	if contentAmount == 0 {
		return 0
	}
	_, exist := cr.RarityTable[contentType]
	if !exist {
		panic(fmt.Sprint("Required amount doesn't exist for: ", contentType, ", ", contentId))
	}
	total, exist := cr.RarityTable[contentType][contentId]
	if !exist {
		panic(fmt.Sprint("Required amount doesn't exist for: ", contentType, ", ", contentId))
	}
	total /= int64(contentAmount)
	if total < 50 || total > 1<<28 {
		panic(fmt.Sprint("Abnormal weight: ", contentType, ", ", contentId, ", ", contentAmount, ": ", total))
	}
	return int32(total)
}

func loadContentRarity(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.ContentRarity = new(ContentRarity)
	gamedata.ContentRarity.RarityTable = map[int32]map[int32]int64{}
	for _, accessory := range gamedata.Accessory {
		// assume we need 9 of every UR accessory for completion
		// this is equal to 9 * 6 copies
		// we also need to level the skill
		// 19 level, each level requires 2 more R compared to the previous starting at 2
		// so 19 * 20 = 380 R worths = around 10.(5) UR
		// then SR is assumed to be 6 times more common, and R 36 times
		// these numbers mean nothing anyway, but we will round the necessary UR to 60 and it should give enough for max out everything
		// this doesn't take into account gold, but it's neglectible, probably
		amount := int32(54 + 6)
		if accessory.RarityType == enum.AccessoryRarityRare {
			amount *= 36
		} else if accessory.RarityType == enum.AccessoryRaritySRare {
			amount *= 6
		}
		amount *= 5 // adjust the chance up at bit
		gamedata.ContentRarity.AddContent(enum.ContentTypeAccessory, accessory.Id, amount)
	}
	// accessory level up item
	for _, levelUpItem := range gamedata.AccessoryLevelUpItem {
		amount := int32(1000)
		if levelUpItem.Rarity == enum.AccessoryRarityRare {
			amount *= 36
		} else if levelUpItem.Rarity == enum.AccessoryRaritySRare {
			amount *= 6
		}
		gamedata.ContentRarity.AddContent(enum.ContentTypeAccessoryLevelUp, levelUpItem.Id, amount)
		// gamedata.ContentRarity.AddContent(item.Gold.ContentType, item.Gold.ContentId, amount*levelUpItem.GameMoney)
	}
	// practice items
	for _, card := range gamedata.Card {
		for _, cell := range card.TrainingTree.TrainingTreeMapping.TrainingTreeCellContents {
			for _, content := range cell.TrainingTreeCellItemSet.Resources {
				gamedata.ContentRarity.AddContent(content.ContentType, content.ContentId, content.ContentAmount)
			}
		}
		// exp and gold
		gamedata.ContentRarity.AddContent(item.Gold.ContentType, item.Gold.ContentId, gamedata.CardLevel[card.CardRarityType].GameMoneyPrefixSum[100])
		gamedata.ContentRarity.AddContent(item.EXP.ContentType, item.EXP.ContentId, gamedata.CardLevel[card.CardRarityType].ExpPrefixSum[100])
	}

	// bond board
	for _, panel := range gamedata.MemberLovePanelCell {
		for _, content := range panel.Resources {
			gamedata.ContentRarity.AddContent(content.ContentType, content.ContentId, content.ContentAmount)
		}
	}
	// for contentType, m := range gamedata.ContentRarity.RarityTable {
	// 	for contentId, contentAmount := range m {
	// 		fmt.Println(contentType, contentId, contentAmount)
	// 	}
	// }
}

func init() {
	addLoadFunc(loadContentRarity)
	addPrequisite(loadContentRarity, loadAccessory)
	addPrequisite(loadContentRarity, loadAccessoryLevelUpItem)
	addPrequisite(loadContentRarity, loadCard)
	addPrequisite(loadContentRarity, loadCardLevel)
	addPrequisite(loadContentRarity, loadMemberLovePanelCell)
}
