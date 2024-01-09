package request

type FetchTowerTopRequest struct {
	TowerId int32 `json:"tower_id"`
}

type ClearedTowerFloorRequest struct {
	TowerId    int32 `json:"tower_id"`
	FloorNo    int32 `json:"floor_no"`
	IsAutoMode bool  `json:"is_auto_mode"`
}

type RecoveryTowerCardUsedRequest struct {
	TowerId       int32   `json:"tower_id"`
	CardMasterIds []int32 `json:"card_master_ids"`
}

type RecoveryTowerCardUsedAllRequest = FetchTowerTopRequest
type FetchTowerRankingRequest = FetchTowerTopRequest
