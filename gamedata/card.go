package gamedata

import (
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
	Id             int32   `xorm:"pk 'id'"`
	MemberMasterId *int32  `xorm:"'member_m_id'"`
	Member         *Member `xorm:"-"`
	// SchoolIdolNo int `xorm:"'school_idol_no'"`
	CardRarityType int32       `xorm:"'card_rarity_type'" enum:"CardRarityType"`
	Rarity         *CardRarity `xorm:"-"`
	Role           int32       `xorm:"'role'"`
	// MemberCardThumbnailAssetPath string
	// AtGacha bool
	// AtEvent bool
	TrainingTreeMasterId *int32        `xorm:"'training_tree_m_id'"` // must be equal to Id
	TrainingTree         *TrainingTree `xorm:"-"`
	// ActiveSkillVoicePath string
	// SpPoint int
	// ExchangeItemId int `xorm:"'exchange_item_id'"`
	// RoleEffectMasterId int `xorm:"'role_effect_master_id'"` // is just the same as role
	PassiveSkillSlot    int32 `xorm:"'passive_skill_slot'"`
	MaxPassiveSkillSlot int32 `xorm:"'max_passive_skill_slot'"`

	// from m_card_grade_up_item
	// map content_id to client.Content
	CardGradeUpItem map[int32](map[int32]client.Content) `xorm:"-"`
}

type CardGradeUpItem struct {
	Grade    int32          `xorm:"'grade'"`
	Resource client.Content `xorm:"extends"`
}

func (card *Card) populate(gamedata *Gamedata) {
	card.Member = gamedata.Member[*card.MemberMasterId]
	card.MemberMasterId = &card.Member.Id
	card.TrainingTree = gamedata.TrainingTree[*card.TrainingTreeMasterId]
	card.TrainingTreeMasterId = &card.TrainingTree.Id
	card.Rarity = gamedata.CardRarity[card.CardRarityType]
	{
		card.CardGradeUpItem = make(map[int32](map[int32]client.Content))
		gradeUps := []CardGradeUpItem{}
		var err error
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_card_grade_up_item").Where("card_id = ?", card.Id).Find(&gradeUps)
		})
		utils.CheckErr(err)
		for _, gradeUp := range gradeUps {
			_, exist := card.CardGradeUpItem[gradeUp.Grade]
			if !exist {
				card.CardGradeUpItem[gradeUp.Grade] = make(map[int32]client.Content)
			}
			card.CardGradeUpItem[gradeUp.Grade][gradeUp.Resource.ContentId] = gradeUp.Resource
		}
	}

	gamedata.CardByMemberId[*card.MemberMasterId] = append(gamedata.CardByMemberId[*card.MemberMasterId], card)
}

func loadCard(gamedata *Gamedata) {
	fmt.Println("Loading Card")
	gamedata.Card = make(map[int32]*Card)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_card").Find(&gamedata.Card)
	})
	utils.CheckErr(err)
	gamedata.CardByMemberId = map[int32][]*Card{}

	for _, card := range gamedata.Card {
		card.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadCard)
	addPrequisite(loadCard, loadCardRarity)
	addPrequisite(loadCard, loadMember)
	addPrequisite(loadCard, loadTrainingTree)
}
