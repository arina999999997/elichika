package client

import (
	"elichika/generic"
)

type PresentHistoryItem struct {
	Content          Content                         `json:"content"`
	PresentRouteType int32                           `json:"present_route_type" enum:"PresentRouteType"`
	PresentRouteId   generic.Nullable[int32]         `json:"present_route_id"`
	ParamServer      generic.Nullable[LocalizedText] `json:"param_server"`
	ParamClient      generic.Nullable[string]        `json:"param_client"`
	HistoryCreatedAt int64                           `json:"history_created_at"`
}
