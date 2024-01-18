package client

import (
	"elichika/generic"
)

type PresentItem struct {
	Id               int32                           `json:"id"`
	Content          Content                         `json:"content"`
	PresentRouteType int32                           `json:"present_route_type" enum:"PresentRouteType"`
	PresentRouteId   generic.Nullable[int32]         `json:"present_route_id"`
	ParamServer      generic.Nullable[LocalizedText] `json:"param_server"` // pointer
	ParamClient      generic.Nullable[string]        `json:"param_client"` // pointer
	PostedAt         int64                           `json:"posted_at"`
	ExpiredAt        generic.Nullable[int64]         `json:"expired_at"`
	IsNew            bool                            `json:"is_new"`
}
