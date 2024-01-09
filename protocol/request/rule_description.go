package request

type SaveReferenceBookRequest struct {
	ReferenceBookId int32 `json:"reference_book_id"`
}

type SaveRuleDescriptionRequest struct {
	RuleDescriptionMasterIds []int32 `json:"rule_description_master_ids"`
}
