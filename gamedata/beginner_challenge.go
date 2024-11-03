package gamedata

import (
	"elichika/client"

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

func (b *BeginnerChallenge) populate(gamedata *Gamedata) {
	var exist bool
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exist, err = session.Table("m_beginner_challenge_complete_reward").Where("challenge_m_id = ?", b.Id).Get(&b.CompleteReward)
	})
	utils.CheckErrMustExist(err, exist)
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exist, err = session.Table("m_beginner_challenge_complete_reward").Where("challenge_m_id = ?", b.Id).Cols("complete_count").Get(&b.CompleteCount)
	})
	utils.CheckErrMustExist(err, exist)
	cellIds := []int32{}
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_challenge_cell").Where("set_m_id = ?", b.CellSetId).OrderBy("display_order").Cols("id").Find(&cellIds)
	})
	utils.CheckErr(err)
	for _, id := range cellIds {
		b.ChallengeCells = append(b.ChallengeCells, gamedata.BeginnerChallengeCell[id])
		gamedata.BeginnerChallengeCell[id].ChallengeId = b.Id
	}
}

func loadBeginnerChallenge(gamedata *Gamedata) {
	fmt.Println("Loading BeginnerChallenge")
	gamedata.BeginnerChallenge = make(map[int32]*BeginnerChallenge)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_beginner_challenge").Find(&gamedata.BeginnerChallenge)
	})
	utils.CheckErr(err)
	for _, beginnerChallenge := range gamedata.BeginnerChallenge {
		beginnerChallenge.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadBeginnerChallenge)
	addPrequisite(loadBeginnerChallenge, loadBeginnerChallengeCell)
}
