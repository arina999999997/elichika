package request

type SaveReferenceBookRequest struct {
	ReferenceBookID int `json:"reference_book_id"`
}

type SaveRuleDescriptionRequest struct {
	RuleDescriptionMasterIDs []int `json:"rule_description_master_ids"`
}
