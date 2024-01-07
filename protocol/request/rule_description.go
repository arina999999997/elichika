package request

type SaveReferenceBookRequest struct {
	ReferenceBookId int `json:"reference_book_id"`
}

type SaveRuleDescriptionRequest struct {
	RuleDescriptionMasterIds []int `json:"rule_description_master_ids"`
}
