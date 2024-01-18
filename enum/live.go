package enum

const (
	LiveFinishStatusSucceeded = 0x00000001
	LiveFinishStatusFailure   = 0x0000000b
	LiveFinishStatusRetired   = 0x0000000c
	LiveFinishStatusSurrender = 0x00000015
)

const (
	LiveUnlockPatternOpen                 = 0x00000001
	LiveUnlockPatternRequiringStoryMain   = 0x00000002 // no longer necessary after the last update
	LiveUnlockPatternClosed               = 0x00000003
	LiveUnlockPatternMemberStory          = 0x00000004
	LiveUnlockPatternCoopOnly             = 0x00000005
	LiveUnlockPatternTowerOnly            = 0x00000006
	LiveUnlockPatternExtra                = 0x00000007
	LiveUnlockPatternPlayableTutorialOnly = 0x00000008
	LiveUnlockPatternStoryOnly            = 0x00000009
	LiveUnlockPatternDaily                = 0x0000000a
)

const (
	LiveTypeManual = 0x00000001
	LiveTypeMv     = 0x00000002
	LiveTypeCoop   = 0x00000003
	LiveTypeTower  = 0x00000004
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
