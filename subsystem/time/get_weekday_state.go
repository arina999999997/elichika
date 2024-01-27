package time

import (
	"elichika/client"
	"elichika/userdata"

	"time"
)

func GetWeekdayState(session *userdata.Session) client.WeekdayState {
	state := client.WeekdayState{}
	year, month, day := session.Time.Date()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, session.Time.Location()).Unix()
	state.Weekday = int32(session.Time.Weekday())
	if state.Weekday == 0 {
		state.Weekday = 7
	}
	state.NextWeekdayAt = tomorrow
	return state
}
