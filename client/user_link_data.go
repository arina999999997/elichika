package client

import (
	"elichika/generic"
)

type UserLinkData struct {
	UserId               int32                                 `json:"user_id"`
	AuthorizationKey     string                                `json:"authorization_key"`
	Name                 LocalizedText                         `json:"name"`
	LastLoginAt          int64                                 `json:"last_login_at"`
	SnsCoin              int32                                 `json:"sns_coin"`
	TermsOfUseVersion    int32                                 `json:"terms_of_use_version"`
	ServiceUserCommonKey generic.Nullable[generic.Array[byte]] `json:"service_user_common_key"`
}
