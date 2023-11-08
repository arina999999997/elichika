package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCardParam struct {
	TrainingContentType int `xorm:"'training_content_type'"`
	Value               int `xorm:"'value'"`
}

type TrainingTreeCardStorySide struct {
	TrainingContentType int `xorm:"'training_content_type'"`
	TrainingContentNo   int `xorm:"'training_content_no'"` // assume to be always 1
	StorySideMID        int `xorm:"'story_side_m_id'"`
}

type TrainingTreeProgressReward struct {
	ActivateNum int           `xorm:"'activate_num'"`
	Reward      model.Content `xorm:"extends"`
}

type TrainingTree struct {
	// from m_training_tree
	ID                       int                  `xorm:"pk 'id'"`
	TrainingTreeMappingMID   *int                 `xorm:"'training_tree_mapping_m_id'"`
	TrainingTreeMapping      *TrainingTreeMapping `xorm:"-"`
	TrainingTreeCardParamMID int                  `xorm:"'training_tree_card_param_m_id'"`
	// from m_training_tree_card_param
	TrainingTreeCardParams []TrainingTreeCardParam `xorm:"-"` // 1 indexed
	// TrainingTreeCardPassiveSkillIncreaseMID int `xorm:"'training_tree_card_passive_skill_increase_m_id;"`
	// from m_training_tree_card_passive_skill_increase
	// basically only differ between level 5 max skill and level 7 max skill, not implemented here
	// TrainingTreeCardPassiveSkillIncrease

	// from m_training_tree_card_story_side
	TrainingTreeCardStorySides map[int]int `xorm:"-"` // map from training_content_type to storySideMID
	// from m_training_tree_card_suit
	SuitMIDs []int `xorm:"-"` // 1 indexed
	// from m_training_tree_card_voice
	NaviActionIDs []int `xorm:"-"` // 1 indexed

	// from m_training_tree_progress_reward
	TrainingTreeProgressRewards []TrainingTreeProgressReward `xorm:"-"`
}

func (tree *TrainingTree) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	tree.TrainingTreeMapping = gamedata.TrainingTreeMapping[*tree.TrainingTreeMappingMID]
	tree.TrainingTreeMappingMID = &tree.TrainingTreeMapping.ID
	{
		err := masterdata_db.Table("m_training_tree_card_param").Where("id = ?", tree.TrainingTreeCardParamMID).
			OrderBy("training_content_no").Find(&tree.TrainingTreeCardParams)
		tree.TrainingTreeCardParams = append([]TrainingTreeCardParam{TrainingTreeCardParam{}}, tree.TrainingTreeCardParams...)
		utils.CheckErr(err)
	}

	{
		tree.TrainingTreeCardStorySides = make(map[int]int)
		stories := []TrainingTreeCardStorySide{}
		err := masterdata_db.Table("m_training_tree_card_story_side").Where("card_m_id = ?", tree.ID).Find(&stories)
		utils.CheckErr(err)
		for _, story := range stories {
			tree.TrainingTreeCardStorySides[story.TrainingContentType] = story.StorySideMID
		}
	}

	{
		err := masterdata_db.Table("m_training_tree_card_suit").Where("card_m_id = ?", tree.ID).
			OrderBy("training_content_no").Cols("suit_m_id").Find(&tree.SuitMIDs)
		utils.CheckErr(err)
		tree.SuitMIDs = append([]int{0}, tree.SuitMIDs...)
	}

	{
		err := masterdata_db.Table("m_training_tree_card_voice").Where("card_m_id = ?", tree.ID).
			OrderBy("training_content_no").Cols("navi_action_id").Find(&tree.NaviActionIDs)
		utils.CheckErr(err)
		tree.NaviActionIDs = append([]int{0}, tree.NaviActionIDs...)
	}

	{
		err := masterdata_db.Table("m_training_tree_progress_reward").Where("card_master_id = ?", tree.ID).
			OrderBy("activate_num").Find(&tree.TrainingTreeProgressRewards)
		utils.CheckErr(err)
	}
}

func loadTrainingTree(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingTree")
	gamedata.TrainingTree = make(map[int]*TrainingTree)
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
