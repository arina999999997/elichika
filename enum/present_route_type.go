package enum

const (
	PresentRouteTypeGacha                             int32 = 0x00000001 // unused for now
	PresentRouteTypeAdminPresent                      int32 = 0x00000002 // handled
	PresentRouteTypeLoginBonus                        int32 = 0x00000003 // handled
	PresentRouteTypeSpecialLoginBonus                 int32 = 0x00000004 // unused for now
	PresentRouteTypeGiftBox                           int32 = 0x00000005 // unused for now
	PresentRouteTypeEventMarathonPointReward          int32 = 0x00000006 // unused for now
	PresentRouteTypeItemFull                          int32 = 0x00000007 // handled
	PresentRouteTypeLoveLevelUp                       int32 = 0x00000008 // handled
	PresentRouteTypeRetryGacha                        int32 = 0x00000009 // unused for now
	PresentRouteTypeShopExchange                      int32 = 0x0000000a // is moved to PresentRouteTypeTrade?
	PresentRouteTypeShopEventExchange                 int32 = 0x0000000b // unused for now
	PresentRouteTypeGpsPresent                        int32 = 0x0000000c // unused for now
	PresentRouteTypeSchoolIdolFestivalIdReward        int32 = 0x0000000d // unused for now
	PresentRouteTypeSerialCode                        int32 = 0x0000000e // unused for now
	PresentRouteTypeEventMarathonRankingReward        int32 = 0x0000000f // unused for now
	PresentRouteTypeMission                           int32 = 0x00000010 // unused for now
	PresentRouteTypeTutorial                          int32 = 0x00000011 // unused for now (for tutorial/beginner reward?)
	PresentRouteTypeStoryMain                         int32 = 0x00000012 // handled
	PresentRouteTypeShopItemFull                      int32 = 0x00000013 // unused (used for paid product only)
	PresentRouteTypeLiveAccessoryItemFull             int32 = 0x00000014 // handled (live doesn't stand for received from live show)
	PresentRouteTypeEventMiningPointRankingReward     int32 = 0x00000015 // unused for now
	PresentRouteTypeEventMiningVoltageRankingReward   int32 = 0x00000016 // unused for now
	PresentRouteTypeEventCoopPointReward              int32 = 0x00000017 // unused for now
	PresentRouteTypeTrade                             int32 = 0x00000018 // handled
	PresentRouteTypeEventCoopGlobalReward             int32 = 0x00000019 // unused for now
	PresentRouteTypeEventCoopPeriodTotalRankingReward int32 = 0x0000001a // unused for now
	PresentRouteTypeEventCoopDailyRankingReward       int32 = 0x0000001b // unused for now
	PresentRouteTypeTowerClearReward                  int32 = 0x0000001c // handled
	PresentRouteTypeTowerProgressReward               int32 = 0x0000001d // handled
	PresentRouteTypeStoryMember                       int32 = 0x0000001e // handled
	PresentRouteTypeSubscriptionDailyReward           int32 = 0x0000001f // unused for now
	PresentRouteTypeSubscriptionContinueReward        int32 = 0x00000020 // unused for now
	PresentRouteTypeStoryLinkage                      int32 = 0x00000021 // unused for now
	PresentRouteTypeStoryLinkageAddtional             int32 = 0x00000022 // unused for now
	PresentRouteTypeExternalMovieReward               int32 = 0x00000023 // unused for now
	PresentRouteTypeMemberGuildSupportReward          int32 = 0x00000024 // handled
	PresentRouteTypeMemberGuildInsideRankingReward    int32 = 0x00000025 // unused for now
	PresentRouteTypeMemberGuildOutsideRankingReward   int32 = 0x00000026 // unused for now
	PresentRouteTypeMemberGuildPointClearReward       int32 = 0x00000027 // unused for now
	PresentRouteTypeTowerBonusLiveRankingReward       int32 = 0x00000028 // unused for now
	PresentRouteTypeSteadyVoltageRankingReward        int32 = 0x00000029 // unused for now
	PresentRouteTypeDebug                             int32 = 0x00000063 // unused for now
)
