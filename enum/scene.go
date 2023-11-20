package enum

// this is experimentaly found so might not be entirely correct
const (
	UnlockSceneStatusInitial  = 1
	UnlockSceneStatusUnlocked = 2
)

const (
	// it's not known clearly what 5 6 7 does
	// There are mention of ShopEventExchange and ShopItemExchange, so maybe once they were used but no longer
	UnlockSceneTypeLessonMenuSelect = 1
	UnlockSceneTypeLiveMusicSelect  = 2
	UnlockSceneTypeAccessoryList    = 3
	UnlockSceneTypeStoryMember      = 4
	// 5
	// 6
	// 7
	UnlockSceneTypeReferenceBookSelect = 8
	// TODO: need to implement member guide properly first, after unlocking member guild, it will be a 2 steps tutorial that first unlock scene tips 3 then 4
	// UnlockSceneTypeMemberGuild = 9

)
