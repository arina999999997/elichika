package request

import (
	"elichika/generic"
)

type SaveRuleDescriptionRequest struct {
	RuleDescriptionMasterIds generic.Array[int32] `json:"rule_description_master_ids"`
}
