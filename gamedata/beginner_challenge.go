package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type BeginnerChallenge struct {
	// from m_beginner_challenge
	Id             int32                    `xorm:"pk 'id'"`
	CellSetId      int32                    `xorm:"'cell_set_id'"`
	ChallengeCells []*BeginnerChallengeCell `xorm:"-"`
	// Title string `xorm:"'title'"`
	// CongratulationsText string `xorm:"'congratulations_text'"`
	StartAt int64 `xorm:"'start_at'"`
	// BackgroundImageAssetPath string `xorm:"background_image_asset_path"`

	// from m_beginner_challenge_complete_reward
	CompleteCount  int32          `xorm:"-"`
	CompleteReward client.Content `xorm:"-"`
}

func (b *BeginnerChallenge) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	exist, err := masterdata_db.Table("m_beginner_challenge_complete_reward").Where("challenge_m_id = ?", b.Id).
		Get(&b.CompleteReward)
	utils.CheckErrMustExist(err, exist)
	exist, err = masterdata_db.Table("m_beginner_challenge_complete_reward").Where("challenge_m_id = ?", b.Id).
		Cols("complete_count").Get(&b.CompleteCount)
	utils.CheckErrMustExist(err, exist)
	cellIds := []int32{}
	err = masterdata_db.Table("m_challenge_cell").Where("set_m_id = ?", b.CellSetId).OrderBy("display_order").
		Cols("id").Find(&cellIds)
	utils.CheckErr(err)
	for _, id := range cellIds {
		b.ChallengeCells = append(b.ChallengeCells, gamedata.BeginnerChallengeCell[id])
		gamedata.BeginnerChallengeCell[id].ChallengeId = b.Id
	}
}

func loadBeginnerChallenge(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading BeginnerChallenge")
	gamedata.BeginnerChallenge = make(map[int32]*BeginnerChallenge)
	err := masterdata_db.Table("m_beginner_challenge").Find(&gamedata.BeginnerChallenge)
	utils.CheckErr(err)
	for _, beginnerChallenge := range gamedata.BeginnerChallenge {
		beginnerChallenge.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadBeginnerChallenge)
	addPrequisite(loadBeginnerChallenge, loadBeginnerChallengeCell)
}
