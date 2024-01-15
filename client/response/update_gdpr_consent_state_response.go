package response

import (
	"elichika/client"
)

type UpdateGdprConsentStateResponse struct {
	UserModel     *client.UserModel            `json:"user_model"`
	ConsentedInfo client.UserGdprConsentedInfo `json:"consented_info"`
}
