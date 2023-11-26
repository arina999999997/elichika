package model

type TrainingTreeCell struct {
	UserID       int   `xorm:"pk 'user_id'" json:"-"`
	CardMasterID int   `xorm:"pk 'card_master_id'" json:"-"`
	CellID       int   `xorm:"pk 'cell_id'" json:"cell_id"`
	ActivatedAt  int64 `json:"activated_at"` // int64 so we don't have Y2K38 problem
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_training_tree_cell"] = TrainingTreeCell{}
}
