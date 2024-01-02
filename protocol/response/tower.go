package response

import (
	"elichika/model"
)

type FetchTowerSelectResponse struct {
	TowerIDs      []int      `json:"tower_ids"`
	UserModelDiff *UserModel `json:"user_model_diff"`
}

type FetchTowerTopResponse struct {
	TowerCardUsedCountRows []model.UserTowerCardUsedCount `json:"tower_card_used_count_rows"`
	IsShowUnlockEffect     bool                           `json:"is_show_unlock_effect"`
	UserModelDiff          *UserModel                     `json:"user_model_diff"`

	// used for bonus lives (ranking one), order will display the small ranking number
	// each bonus live voltage is an array of number of floor length that store the score to them
	Order                *int  `json:"order"`
	EachBonusLiveVoltage []int `json:"each_bonus_live_voltage"`
}

type ClearedTowerFloorResponse struct {
	IsShowUnlockEffect bool       `json:"is_show_unlock_effect"`
	UserModelDiff      *UserModel `json:"user_model_diff"`
}

type RecoveryTowerCardUsedResponse struct {
	TowerCardUsedCountRows []model.UserTowerCardUsedCount `json:"tower_card_used_count_rows"`
	UserModelDiff          *UserModel                     `json:"user_model_diff"`
}

type RecoveryTowerCardUsedAllResponse struct {
	UserModelDiff *UserModel `json:"user_model_diff"`
}

type TowerRankingUser = RankingUser

type TowerRankingCell struct {
	Order            int              `json:"order"`
	SumVoltage       int              `json:"sum_voltage"`
	TowerRankingUser TowerRankingUser `json:"tower_ranking_user"`
}

type TowerRankingBorderMasterRow struct {
	RankingType  int `json:"ranking_type"` //  EventCommonRankingType
	UpperRank    int `json:"upper_rank"`
	LowerRank    int `json:"lower_rank"`
	DisplayOrder int `json:"display_order"`
}

type TowerRankingBorderInfo struct {
	RankingBorderVoltage   int                         `json:"ranking_border_voltage"`
	RankingBorderMasterRow TowerRankingBorderMasterRow `json:"ranking_border_master_row"`
}

type FetchTowerRankingResponse struct {
	TopRankingCells    []TowerRankingCell       `json:"top_ranking_cells"`
	MyRankingCells     []TowerRankingCell       `json:"my_ranking_cells"`
	FriendRankingCells []TowerRankingCell       `json:"friend_ranking_cells"`
	RankingBorderInfo  []TowerRankingBorderInfo `json:"ranking_border_info"`
	MyOrder            int                      `json:"my_order"`
}
