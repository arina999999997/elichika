package client

import (
	"elichika/generic"
)

type PresentItem struct {
	Id               int32                           `xorm:"pk" json:"id"`
	Content          Content                         `xorm:"extends" json:"content"`
	PresentRouteType int32                           `json:"present_route_type" enum:"PresentRouteType"`
	PresentRouteId   generic.Nullable[int32]         `xorm:"json" json:"present_route_id"`
	ParamServer      generic.Nullable[LocalizedText] `xorm:"json" json:"param_server"` // pointer
	ParamClient      generic.Nullable[string]        `xorm:"json" json:"param_client"` // pointer
	PostedAt         int64                           `json:"posted_at"`
	ExpiredAt        generic.Nullable[int64]         `xorm:"json" json:"expired_at"`
	IsNew            bool                            `json:"is_new"`
}
