package gamedata

import (
	"elichika/dictionary"
	// "elichika/model"
	"elichika/client"
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
	Id             int     `xorm:"pk 'id'"`
	MemberMasterId *int    `xorm:"'member_m_id'"`
	Member         *Member `xorm:"-"`
	// SchoolIdolNo int `xorm:"'school_idol_no'"`
	CardRarityType int `xorm:"'card_rarity_type'"`
	Role           int `xorm:"'role'"`
	// MemberCardThumbnailAssetPath string
	// AtGacha bool
	// AtEvent bool
	TrainingTreeMasterId *int          `xorm:"'training_tree_m_id'"` // must be equal to Id
	TrainingTree         *TrainingTree `xorm:"-"`
	// ActiveSkillVoicePath string
	// SpPoint int
	// ExchangeItemId int `xorm:"'exchange_item_id'"`
	// RoleEffectMasterId int `xorm:"'role_effect_master_id'"` // is just the same as role
	PassiveSkillSlot    int `xorm:"'passive_skill_slot'"`
	MaxPassiveSkillSlot int `xorm:"'max_passive_skill_slot'"`

	// from m_card_grade_up_item
	// map content_id to client.Content
	CardGradeUpItem map[int](map[int32]client.Content) `xorm:"-"`
}

type CardGradeUpItem struct {
	Grade    int            `xorm:"'grade'"`
	Resource client.Content `xorm:"extends"`
}

func (card *Card) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	card.Member = gamedata.Member[*card.MemberMasterId]
	card.MemberMasterId = &card.Member.Id
	card.TrainingTree = gamedata.TrainingTree[*card.TrainingTreeMasterId]
	card.TrainingTreeMasterId = &card.TrainingTree.Id

	{
		card.CardGradeUpItem = make(map[int](map[int32]client.Content))
		gradeUps := []CardGradeUpItem{}
		err := masterdata_db.Table("m_card_grade_up_item").Where("card_id = ?", card.Id).Find(&gradeUps)
		utils.CheckErr(err)
		for _, gradeUp := range gradeUps {
			_, exist := card.CardGradeUpItem[gradeUp.Grade]
			if !exist {
				card.CardGradeUpItem[gradeUp.Grade] = make(map[int32]client.Content)
			}
			card.CardGradeUpItem[gradeUp.Grade][gradeUp.Resource.ContentId] = gradeUp.Resource
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
