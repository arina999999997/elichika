package client

import (
	"elichika/generic"
)

type AddedCardResult struct {
	CardMasterId         int32                     `json:"card_master_id"`
	Level                int32                     `json:"level"`
	BeforeGrade          int32                     `json:"before_grade"`
	AfterGrade           int32                     `json:"after_grade"`
	Content              generic.Nullable[Content] `json:"content"` // pointer
	LimitExceeded        bool                      `json:"limit_exceeded"`
	BeforeLoveLevelLimit int32                     `json:"before_love_level_limit"`
	AfterLoveLevelLimit  int32                     `json:"after_love_level_limit"`
}
