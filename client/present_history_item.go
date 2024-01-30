package client

import (
	"elichika/generic"
)

type PresentHistoryItem struct {
	Content          Content                         `xorm:"extends" json:"content"`
	PresentRouteType int32                           `json:"present_route_type" enum:"PresentRouteType"`
	PresentRouteId   generic.Nullable[int32]         `xorm:"json" json:"present_route_id"`
	ParamServer      generic.Nullable[LocalizedText] `xorm:"json" json:"param_server"`
	ParamClient      generic.Nullable[string]        `xorm:"json" json:"param_client"`
	HistoryCreatedAt int64                           `json:"history_created_at"`
}
