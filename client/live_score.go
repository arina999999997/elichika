package client

import (
	"elichika/generic"
)

type LiveScore struct {
	StartInfo                  LiveSpotInfo                             `json:"start_info"`
	FinishInfo                 LiveSpotInfo                             `json:"finish_info"`
	ResultDict                 generic.Dictionary[int32, LiveNoteScore] `json:"result_dict"`
	WaveStatDict               generic.Dictionary[int32, LiveWaveStat]  `json:"wave_stat_dict"`
	TurnStatDict               generic.Dictionary[int32, LiveTurnStat]  `json:"turn_stat_dict"`
	CardStatDict               generic.Dictionary[int32, LiveCardStat]  `json:"card_stat_dict"`
	TargetScore                int32                                    `json:"target_score"`
	CurrentScore               int32                                    `json:"current_score"`
	ComboCount                 int32                                    `json:"combo_count"`
	ChangeSquadCount           int32                                    `json:"change_squad_count"`
	HighestComboCount          int32                                    `json:"highest_combo_count"`
	RemainingStamina           int32                                    `json:"remaining_stamina"`
	IsPerfectLive              bool                                     `json:"is_perfect_live"`
	IsPerfectFullCombo         bool                                     `json:"is_perfect_full_combo"`
	UseVoltageActiveSkillCount int32                                    `json:"use_voltage_active_skill_count"`
	UseHealActiveSkillCount    int32                                    `json:"use_heal_active_skill_count"`
	UseDebufActiveSkillCount   int32                                    `json:"use_debuf_active_skill_count"`
	UseBufActiveSkillCount     int32                                    `json:"use_buf_active_skill_count"`
	UseSpSkillCount            int32                                    `json:"use_sp_skill_count"`
	CompleteAppealChanceCount  int32                                    `json:"complete_appeal_chance_count"`
	TriggeredCriticalCount     int32                                    `json:"triggered_critical_count"`
	LivePower                  int32                                    `json:"live_power"`
	SpSkillScoreList           generic.List[int32]                      `json:"sp_skill_score_list"`
}
