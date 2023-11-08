package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

/*
Assume the following result in the DB:
- SELECT * from m_card WHERE training_tree_m_id != id -> 0 record.
*/
type Card struct {
	// from m_card
	ID             int     `xorm:"pk 'id'"`
	MemberMasterID *int    `xorm:"'member_m_id'"`
	Member         *Member `xorm:"-"`
	// SchoolIdolNo int `xorm:"'school_idol_no'"`
	CardRarityType int `xorm:"'card_rarity_type'"`
	Role           int `xorm:"'role'"`
	// MemberCardThumbnailAssetPath string
	// AtGacha bool
	// AtEvent bool
	TrainingTreeMasterID *int          `xorm:"'training_tree_m_id'"` // must be equal to ID
	TrainingTree         *TrainingTree `xorm:"-"`
	// ActiveSkillVoicePath string
	// SpPoint int
	// ExchangeItemID int `xorm:"'exchange_item_id'"`
	// RoleEffectMasterID int `xorm:"'role_effect_master_id'"` // is just the same as role
	PassiveSkillSlot    int `xorm:"'passive_skill_slot'"`
	MaxPassiveSkillSlot int `xorm:"'max_passive_skill_slot'"`

	// from m_card_grade_up_item
	// ma content_id to model.Content
	CardGradeUpItem map[int](map[int]model.Content) `xorm:"-"`
}

type CardGradeUpItem struct {
	Grade    int           `xorm:"'grade'"`
	Resource model.Content `xorm:"extends"`
}

func (card *Card) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	card.Member = gamedata.Member[*card.MemberMasterID]
	card.MemberMasterID = &card.Member.ID
	card.TrainingTree = gamedata.TrainingTree[*card.TrainingTreeMasterID]
	card.TrainingTreeMasterID = &card.TrainingTree.ID

	{
		card.CardGradeUpItem = make(map[int](map[int]model.Content))
		gradeUps := []CardGradeUpItem{}
		err := masterdata_db.Table("m_card_grade_up_item").Where("card_id = ?", card.ID).Find(&gradeUps)
		utils.CheckErr(err)
		for _, gradeUp := range gradeUps {
			_, exists := card.CardGradeUpItem[gradeUp.Grade]
			if !exists {
				card.CardGradeUpItem[gradeUp.Grade] = make(map[int]model.Content)
			}
			card.CardGradeUpItem[gradeUp.Grade][gradeUp.Resource.ContentID] = gradeUp.Resource
		}
	}
}

func loadCard(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Card")
	gamedata.Card = make(map[int]*Card)
	err := masterdata_db.Table("m_card").Find(&gamedata.Card)
	utils.CheckErr(err)

	for _, card := range gamedata.Card {
		card.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadCard)
	addPrequisite(loadCard, loadMember)
	addPrequisite(loadCard, loadTrainingTree)
}
