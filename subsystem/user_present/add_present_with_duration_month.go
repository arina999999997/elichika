package user_present

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

// duration is expressed in month, this is just a helper and it might not even be necessary
func AddPresentWithDurationMonth(session *userdata.Session, present client.PresentItem, monthCount int) {
	endPoint := session.Time.AddDate(0, monthCount, 0)
	present.ExpiredAt = generic.NewNullable(endPoint.Unix())
	AddPresent(session, present)
}
