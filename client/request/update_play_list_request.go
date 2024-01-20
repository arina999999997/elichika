package request

type UpdatePlayListRequest struct {
	LiveMasterId int32 `json:"live_master_id"`
	GroupNum     int32 `json:"group_num"`
	IsSet        bool  `json:"is_set"`
}
