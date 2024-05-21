package enum

const (
	RecoverableExceptionTypeGachaOutOfTerm                          int32 = 0x00000001
	RecoverableExceptionTypeGachaDailyLimit                         int32 = 0x00000002
	RecoverableExceptionTypeGachaRetryOutOfTerm                     int32 = 0x00000003
	RecoverableExceptionTypeShopItemOutOfTerm                       int32 = 0x00000004
	RecoverableExceptionTypeShopEventOutOfTerm                      int32 = 0x00000005
	RecoverableExceptionTypeShopItemAllExpired                      int32 = 0x00000006
	RecoverableExceptionTypeShopEventAllExpired                     int32 = 0x00000007
	RecoverableExceptionTypeLpLimitExceeded                         int32 = 0x00000008
	RecoverableExceptionTypeApLimitExceeded                         int32 = 0x00000009
	RecoverableExceptionTypeNameContainsNgWord                      int32 = 0x0000000a
	RecoverableExceptionTypeNicknameContainsNgWord                  int32 = 0x0000000b
	RecoverableExceptionTypeMessageContainsNgWord                   int32 = 0x0000000c
	RecoverableExceptionTypeNotBillingSupported                     int32 = 0x0000000d
	RecoverableExceptionTypeShopTopBannerExpired                    int32 = 0x0000000e
	RecoverableExceptionTypeBillingNotGetProductList                int32 = 0x0000000f
	RecoverableExceptionTypeBillingIsStillPurchasing                int32 = 0x00000010
	RecoverableExceptionTypeBillingSuccessResume                    int32 = 0x00000011
	RecoverableExceptionTypeBillingPurchaseCancel                   int32 = 0x00000012
	RecoverableExceptionTypeBillingFailurePlatformConnectShopNormal int32 = 0x00000013
	RecoverableExceptionTypeBillingFailurePurchaseShopNormal        int32 = 0x00000014
	RecoverableExceptionTypeEventMarathonOutOfDate                  int32 = 0x00000015
	RecoverableExceptionTypeBillingExpiredProduct                   int32 = 0x00000016
	RecoverableExceptionTypeCommonNgWord                            int32 = 0x00000017
	RecoverableExceptionTypeBillingNotBuyableProduct                int32 = 0x00000018
	RecoverableExceptionTypeBirthdayValidateErrorMonth              int32 = 0x00000019
	RecoverableExceptionTypeBirthdayValidateErrorDay                int32 = 0x0000001a
	RecoverableExceptionTypeGpsPresentOutOfTermError                int32 = 0x0000001b
	RecoverableExceptionTypeBillingFailurePlatformConnectShopPack   int32 = 0x0000001c
	RecoverableExceptionTypeBillingFailurePurchaseShopPack          int32 = 0x0000001d
	RecoverableExceptionTypeMissionChallengeOutOfDete               int32 = 0x0000001e
	RecoverableExceptionTypeEventMiningOutOfDate                    int32 = 0x0000001f
	RecoverableExceptionTypeShopSuitOutOfTerm                       int32 = 0x00000020
	RecoverableExceptionTypeShopOther1OutOfTerm                     int32 = 0x00000021
	RecoverableExceptionTypeShopOther2OutOfTerm                     int32 = 0x00000022
	RecoverableExceptionTypeShopOther3OutOfTerm                     int32 = 0x00000023
	RecoverableExceptionTypeDailyTheaterOutOfTerm                   int32 = 0x00000024
	RecoverableExceptionTypeWsnetRoomClosed                         int32 = 0x00000025
	RecoverableExceptionTypeWsnetRoomMemberOver                     int32 = 0x00000026
	RecoverableExceptionTypeWsnetUnknown                            int32 = 0x00000027
	RecoverableExceptionTypeEventCoopOutOfDate                      int32 = 0x00000028
	RecoverableExceptionTypeStoryEventHistoryOutOfTerm              int32 = 0x00000029
	RecoverableExceptionTypeDailyLiveOutOfDate                      int32 = 0x0000002a
	RecoverableExceptionTypeDailyLiveNotFillCount                   int32 = 0x0000002b
	RecoverableExceptionTypeEventCoopRoomMemberShortage             int32 = 0x0000002c
	RecoverableExceptionTypeTradeOutOfTerm                          int32 = 0x0000002d
	RecoverableExceptionTypeTowerOutOfTerm                          int32 = 0x0000002e
	RecoverableExceptionTypeTowerSelectOutOfTerm                    int32 = 0x0000002f
	RecoverableExceptionTypeMemberGuildTransferOutOfTerm            int32 = 0x00000030
	RecoverableExceptionTypeMemberGuildTransferAlready              int32 = 0x00000031
	RecoverableExceptionTypeDailyTheaterLikeLimit                   int32 = 0x00000032
	RecoverableExceptionTypeDailyTheaterArchiveOutOfTerm            int32 = 0x00000033
	RecoverableExceptionTypeTowerRankingAggregationTerm             int32 = 0x00000034
	RecoverableExceptionTypeBillingPurchaseReject                   int32 = 0x00000063
	RecoverableExceptionTypeSubscriptionExpiredProduct              int32 = 0x00000064
	RecoverableExceptionTypeSubscriptionPendingPayment              int32 = 0x00000065
	RecoverableExceptionTypeSubscriptionLinkedOtherPlayer           int32 = 0x00000066
	RecoverableExceptionTypeSubscriptionDuplicateSubscribe          int32 = 0x00000067
	RecoverableExceptionTypeSubscriptionEndedTrial                  int32 = 0x00000068
	RecoverableExceptionTypeSubscriptionRestoreError                int32 = 0x0000006a
	RecoverableExceptionTypeSubscriptionRestoreLinkedOtherPlayer    int32 = 0x0000006b
	RecoverableExceptionTypeStoryLinkageChapterExpired              int32 = 0x0000006c
	RecoverableExceptionTypeExternalMovieOutOfTerm                  int32 = 0x0000006d
	RecoverableExceptionTypeExternalMovieRewardOutOfTerm            int32 = 0x0000006e
	RecoverableExceptionTypeDailyLiveOutOfTerm                      int32 = 0x0000006f
	RecoverableExceptionTypeDailyLiveOutOfTermWithoutTransition     int32 = 0x00000070
	RecoverableExceptionTypeUserAccountDeleted                      int32 = 0x00000071
	RecoverableExceptionTypeReserve1                                int32 = 0x000000fb
	RecoverableExceptionTypeReserve2                                int32 = 0x000000fc
	RecoverableExceptionTypeReserve3                                int32 = 0x000000fd
	RecoverableExceptionTypeReserve4                                int32 = 0x000000fe
	RecoverableExceptionTypeReserve5                                int32 = 0x000000ff
)
