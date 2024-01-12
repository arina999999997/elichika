package client

type DoPowerUp struct {
	DoLevelUp        bool `json:"do_level_up"`
	DoGradeUp        bool `json:"do_grade_up"`
	DoAddSkill       bool `json:"do_add_skill"`
	DoSkillProcessed bool `json:"do_skill_processed"`
	DoSkillLevelUp   bool `json:"do_skill_level_up"`
}
