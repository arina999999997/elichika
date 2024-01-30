package user_present

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

// duration is expressed in seconds
func AddPresentWithDuration(session *userdata.Session, present client.PresentItem, duration int64) {
	present.ExpiredAt = generic.NewNullable(session.Time.Unix() + duration)
	AddPresent(session, present)
}
