package user_status

import (
	"elichika/userdata"
)

// lp can be negative
func AddUserLp(session *userdata.Session, lp int32) {
	// live_point_full_at and live_point_broken work like this:
	// - live_point_full_at specify when the live_point will recover to full
	// - if live_point_full_at is later than the current time, then live_point_broken is set to whatever it should be when it's full
	// - if life_point_full_at is earlier or equal to the current time, then live_point_broken would be the amount that of LP the user have

	maxLp := session.Gamedata.UserRank[session.UserStatus.Rank].MaxLp
	currentLp := session.UserStatus.LivePointBroken
	const LivePointRecoverlyAt int32 = 240
	if session.Time.Unix() < session.UserStatus.LivePointFullAt { // already full
		// calculate the current LP using the recovery rate, this is defined using m_constant LivePointRecoverlyAt
		timeLeft := int32(session.UserStatus.LivePointFullAt - session.Time.Unix())

		toRecover := timeLeft / LivePointRecoverlyAt
		if timeLeft%LivePointRecoverlyAt != 0 {
			toRecover++
		}
		currentLp = session.UserStatus.LivePointBroken - toRecover
	}
	currentLp += lp
	if currentLp >= maxLp {
		session.UserStatus.LivePointBroken = currentLp
		session.UserStatus.LivePointFullAt = session.Time.Unix()
	} else {
		session.UserStatus.LivePointBroken = maxLp
		session.UserStatus.LivePointFullAt -= int64(lp * LivePointRecoverlyAt)
	}
}
