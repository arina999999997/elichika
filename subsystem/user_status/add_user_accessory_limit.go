package user_status

import (
	"elichika/userdata"
)

func AddUserAccessoryLimit(session *userdata.Session, accessoryLimit int32) {
	session.UserStatus.AccessoryBoxAdditional += accessoryLimit
}
