package request

import (
	"elichika/generic"
)

type ActivateTrainingTreeCellRequest struct {
	CardMasterId  int32               `json:"card_master_id"`                      // is actually named _CardMasterId
	CellMasterIds generic.List[int32] `json:"cell_master_ids"`                     // is actually named _CellMasterIds
	PayType       int32               `json:"pay_type" enum:"TrainingTreePayType"` // is actually named _PayType
}
