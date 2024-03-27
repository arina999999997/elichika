package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCardParam struct {
	TrainingContentType int32 `xorm:"'training_content_type'"`
	Value               int32 `xorm:"'value'"`
}

type TrainingTreeCardStorySide struct {
	TrainingContentType int32 `xorm:"'training_content_type'"`
	TrainingContentNo   int32 `xorm:"'training_content_no'"` // assume to be always 1
	StorySideMId        int32 `xorm:"'story_side_m_id'"`
}

type TrainingTreeProgressReward struct {
	ActivateNum int            `xorm:"'activate_num'"`
	Reward      client.Content `xorm:"extends"`
}

type TrainingTree struct {
	// from m_training_tree
	Id                       int32                `xorm:"pk 'id'"`
	TrainingTreeMappingMId   *int32               `xorm:"'training_tree_mapping_m_id'"`
	TrainingTreeMapping      *TrainingTreeMapping `xorm:"-"`
	TrainingTreeCardParamMId int32                `xorm:"'training_tree_card_param_m_id'"`
	// from m_training_tree_card_param
	TrainingTreeCardParams []TrainingTreeCardParam `xorm:"-"` // 1 indexed
	// TrainingTreeCardPassiveSkillIncreaseMId int `xorm:"'training_tree_card_passive_skill_increase_m_id;"`
	// from m_training_tree_card_passive_skill_increase
	// basically only differ between level 5 max skill and level 7 max skill, not implemented here
	// TrainingTreeCardPassiveSkillIncrease

	// from m_training_tree_card_story_side
	TrainingTreeCardStorySides map[int32]int32 `xorm:"-"` // map from training_content_type to storySideMId
	// from m_training_tree_card_suit
	SuitMIds []int32 `xorm:"-"` // 1 indexed
	// from m_training_tree_card_voice
	NaviActionIds []int32 `xorm:"-"` // 1 indexed

	// from m_training_tree_progress_reward
	TrainingTreeProgressRewards []TrainingTreeProgressReward `xorm:"-"`
}

func (tree *TrainingTree) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	tree.TrainingTreeMapping = gamedata.TrainingTreeMapping[*tree.TrainingTreeMappingMId]
	tree.TrainingTreeMappingMId = &tree.TrainingTreeMapping.Id
	{
		err := masterdata_db.Table("m_training_tree_card_param").Where("id = ?", tree.TrainingTreeCardParamMId).
			OrderBy("training_content_no").Find(&tree.TrainingTreeCardParams)
		tree.TrainingTreeCardParams = append([]TrainingTreeCardParam{TrainingTreeCardParam{}}, tree.TrainingTreeCardParams...)
		utils.CheckErr(err)
	}

	{
		tree.TrainingTreeCardStorySides = make(map[int32]int32)
		stories := []TrainingTreeCardStorySide{}
		err := masterdata_db.Table("m_training_tree_card_story_side").Where("card_m_id = ?", tree.Id).Find(&stories)
		utils.CheckErr(err)
		for _, story := range stories {
			tree.TrainingTreeCardStorySides[story.TrainingContentType] = story.StorySideMId
		}
	}

	{
		err := masterdata_db.Table("m_training_tree_card_suit").Where("card_m_id = ?", tree.Id).
			OrderBy("training_content_no").Cols("suit_m_id").Find(&tree.SuitMIds)
		utils.CheckErr(err)
		tree.SuitMIds = append([]int32{0}, tree.SuitMIds...)
	}

	{
		err := masterdata_db.Table("m_training_tree_card_voice").Where("card_m_id = ?", tree.Id).
			OrderBy("training_content_no").Cols("navi_action_id").Find(&tree.NaviActionIds)
		utils.CheckErr(err)
		tree.NaviActionIds = append([]int32{0}, tree.NaviActionIds...)
	}

	{
		err := masterdata_db.Table("m_training_tree_progress_reward").Where("card_master_id = ?", tree.Id).
			OrderBy("activate_num").Find(&tree.TrainingTreeProgressRewards)
		utils.CheckErr(err)
	}
}

func loadTrainingTree(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingTree")
	gamedata.TrainingTree = make(map[int32]*TrainingTree)
	err := masterdata_db.Table("m_training_tree").Find(&gamedata.TrainingTree)
	utils.CheckErr(err)
	for _, tree := range gamedata.TrainingTree {
		tree.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadTrainingTree)
	addPrequisite(loadTrainingTree, loadTrainingTreeMapping)
}
