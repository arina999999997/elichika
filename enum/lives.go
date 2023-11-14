package enum

const (
	LiveFinishStatusCleared  = 1
	LiveFinishStatusFailure  = 11
	LiveFinishStatusUserExit = 12
)

var (
	// beginner, intermediate, advanced, expert (advanced+), challenge
	// 40 is also used in network (not in db), but not sure what it represent
	// TODO: use a map instead of array for this
	LiveDifficultyTypes = [5]int{10, 20, 30, 35, 37}
	LiveDifficultyIndex = map[int]int{
		10: 0,
		20: 1,
		30: 2,
		35: 3,
		37: 4,
	}
)
