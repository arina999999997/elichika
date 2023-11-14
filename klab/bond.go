package klab

func CenterBondGainBasedOnBondGain(gain int) int {
	// https://suyo.be/sifas/wiki/gameplay/live-rewards
	// this doesn't seem to be in DB, so it's probably server sided
	// the amount of bond and LP isn't consistent, because story stage give less bond
	// but the amount of center bond is consistent with the amount of bond per card gain
	// bond for double center is just bond for center / 2, without rounding
	// ignoring the story mode, then the extra center bond has a nice relationship with the lp used:
	// 4 = 0.4 * 10
	// 6 = 0.5 * 12
	// 9 = 0.6 * 15
	// 14 = 0.7 * 20
	if gain < 9 {
		return 0
	} else if gain == 9 {
		return 3
	} else if gain <= 13 {
		return 4
	} else if gain <= 16 {
		return 6
	} else if gain <= 24 {
		return 9
	} else {
		return 14
	}
}
