package client

type UserLinkDataBeforeLogin struct {
	LinkedData  UserLinkData    `json:"linked_data"`
	CurrentData CurrentUserData `json:"current_data"`
}
