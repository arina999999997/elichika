package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

func loadMemberLoveLevel(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading member love level")
	err := masterdata_db.Table("m_member_love_level").OrderBy("love_level").Cols("love_point").Find(&gamedata.MemberLoveLevelLovePoint)
	utils.CheckErr(err)
	gamedata.MemberLoveLevelCount = len(gamedata.MemberLoveLevelLovePoint)
	gamedata.MemberLoveLevelLovePoint = append([]int{0}, gamedata.MemberLoveLevelLovePoint...)
}

func init() {
	addLoadFunc(loadMemberLoveLevel)
}

func (gamedata *Gamedata) LoveLevelFromLovePoint(lovePoint int) int {
	low := 1
	high := gamedata.MemberLoveLevelCount
	mid := 0
	res := 0
	for low <= high {
		mid = (low + high) / 2
		if gamedata.MemberLoveLevelLovePoint[mid] <= lovePoint {
			res = mid
			low = mid + 1		
		} else {
			high = mid - 1
		}
	}
	return res
}