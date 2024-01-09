package client

type UserRuleDescription struct {
	// TODO(refactor): This field doesn't exist in client side, but it's necessary in the database
	RuleDescriptionId int32 `xorm:"pk 'rule_description_id'" json:"-"`
	DisplayStatus     int32 `xorm:"'display_status'" json:"display_status" enum:"RuleDescriptionDisplayStatus"`
}

func (urd *UserRuleDescription) Id() int64 {
	return int64(urd.RuleDescriptionId)
}
func (urd *UserRuleDescription) SetId(id int64) {
	urd.RuleDescriptionId = int32(id)
}
