package model

type LiveFinishCard struct {
	CardMasterId        int `json:"-"`
	GotVoltage          int `json:"got_voltage"`
	SkillTriggeredCount int `json:"skill_triggered_count"`
	AppealCount         int `json:"appeal_count"`
}

func (obj *LiveFinishCard) SetId(id int64) {
	obj.CardMasterId = int(id)
}
