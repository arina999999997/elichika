package request

type FetchTowerTopRequest struct {
	TowerId int `json:"tower_id"`
}

type ClearedTowerFloorRequest struct {
	TowerId    int  `json:"tower_id"`
	FloorNo    int  `json:"floor_no"`
	IsAutoMode bool `json:"is_auto_mode"`
}

type RecoveryTowerCardUsedRequest struct {
	TowerId       int   `json:"tower_id"`
	CardMasterIds []int `json:"card_master_ids"`
}

type RecoveryTowerCardUsedAllRequest = FetchTowerTopRequest
type FetchTowerRankingRequest = FetchTowerTopRequest
