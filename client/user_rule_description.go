package client

type UserRuleDescription struct {
	DisplayStatus int32 `xorm:"'display_status'" json:"display_status" enum:"RuleDescriptionDisplayStatus"`
}
