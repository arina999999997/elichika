package enum

const (
	LiveUnlockPatternOpen                 int32 = 0x00000001
	LiveUnlockPatternRequiringStoryMain   int32 = 0x00000002 // no longer necessary after the last update
	LiveUnlockPatternClosed               int32 = 0x00000003
	LiveUnlockPatternMemberStory          int32 = 0x00000004
	LiveUnlockPatternCoopOnly             int32 = 0x00000005
	LiveUnlockPatternTowerOnly            int32 = 0x00000006
	LiveUnlockPatternExtra                int32 = 0x00000007
	LiveUnlockPatternPlayableTutorialOnly int32 = 0x00000008
	LiveUnlockPatternStoryOnly            int32 = 0x00000009
	LiveUnlockPatternDaily                int32 = 0x0000000a
)
