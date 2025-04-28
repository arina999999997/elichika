package user_training_tree

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/userdata"
	"elichika/utils"

	"xorm.io/xorm"
)

// rewrite the user training tree:
// - fetch the current tree
// - uncompress and then compress it
// - delete irrelevant cells

type FixUserCardTrainingTreeCell struct {
	UserId       int32 `xorm:"pk 'user_id'"`
	CardMasterId int32 `xorm:"pk 'card_master_id'"`
	CellId       int32 `xorm:"pk 'cell_id'"`
	ActivatedAt  int64 `xorm:"'activated_at'"`
}

var fixInsertCells = []FixUserCardTrainingTreeCell{}

func fixInserts(session *xorm.Session) {
	err := session.Begin()
	utils.CheckErr(err)
	_, err = session.Table("u_card_training_tree_cell").Where("TRUE").Delete()
	utils.CheckErr(err)
	n := len(fixInsertCells)
	if n == 0 {
		return
	}
	for i := 0; i < n; i += 1024 {
		end := i + 1024
		if end > n {
			end = n
		}
		_, err = session.Table("u_card_training_tree_cell").Insert(fixInsertCells[i:end])
		utils.CheckErr(err)
	}
	err = session.Commit()
	utils.CheckErr(err)
}

func fixCard(session *xorm.Session, cells []FixUserCardTrainingTreeCell) {
	storedCells := []client.UserCardTrainingTreeCell{}
	for _, cell := range cells {
		storedCells = append(storedCells, client.UserCardTrainingTreeCell{
			CellId:      cell.CellId,
			ActivatedAt: cell.ActivatedAt,
		})
	}
	userId := cells[0].UserId
	cardMasterId := cells[0].CardMasterId

	// uncompressed := gamedata.Instance.TrainingTree[cardMasterId].Design().Uncompress(storedCells)
	recompressed := gamedata.Instance.TrainingTree[cardMasterId].Design().Compress(storedCells)
	for _, cell := range recompressed {
		fixInsertCells = append(fixInsertCells, FixUserCardTrainingTreeCell{
			UserId:       userId,
			CardMasterId: cardMasterId,
			CellId:       cell.CellId,
			ActivatedAt:  cell.ActivatedAt,
		})
	}
}

func FixUsersTrainingTrees() {
	session := userdata.Engine.NewSession()
	if session == nil {
		panic("can't get session")
	}

	// this will requires some memory to run
	storedCells := []FixUserCardTrainingTreeCell{}
	err := session.Table("u_card_training_tree_cell").OrderBy("user_id").
		OrderBy("card_master_id").OrderBy("cell_id").Find(&storedCells)
	utils.CheckErr(err)
	n := len(storedCells)
	i := 0
	for i < n {
		j := i + 1
		for j < n {
			if storedCells[j-1].UserId != storedCells[j].UserId {
				break
			}
			if storedCells[j-1].CardMasterId != storedCells[j].CardMasterId {
				break
			}
			j++
		}
		// storedCells[i:j] store the info for one single card
		fixCard(session, storedCells[i:j])
		i = j
	}
	fixInserts(session)

}
