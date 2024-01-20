package router

import (
	"elichika/handler"
	"elichika/handler/live"
	"elichika/middleware"
	"elichika/webui"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/static", "static")
	{
		api := r.Group("/", middleware.Common)
		api.POST("/asset/getPackUrl", handler.GetPackUrl)

		api.POST("/billing/fetchBillingHistory", handler.BillingHistory)
		api.POST("/billing/updateSubscription", handler.UpdateSubscription)

		api.POST("/bootstrap/fetchBootstrap", handler.FetchBootstrap)
		api.POST("/bootstrap/getClearedPlatformAchievement", handler.GetClearedPlatformAchievement)

		api.POST("/card/changeFavorite", handler.ChangeFavorite)
		api.POST("/card/changeIsAwakeningImage", handler.ChangeIsAwakeningImage)
		api.POST("/card/getOtherUserCard", handler.GetOtherUserCard)
		api.POST("/card/updateCardNewFlag", handler.UpdateCardNewFlag)

		api.POST("/communicationMember/fetchCommunicationMemberDetail", handler.FetchCommunicationMemberDetail)
		api.POST("/communicationMember/finishUserStoryMember", handler.FinishUserStoryMember)
		api.POST("/communicationMember/finishUserStorySide", handler.FinishUserStorySide)
		api.POST("/communicationMember/setTheme", handler.SetTheme)
		api.POST("/communicationMember/setFavoriteMember", handler.SetFavoriteMember)
		api.POST("/communicationMember/updateUserCommunicationMemberDetailBadge", handler.UpdateUserCommunicationMemberDetailBadge)
		api.POST("/communicationMember/updateUserLiveDifficultyNewFlag", handler.UpdateUserLiveDifficultyNewFlag)

		api.POST("/emblem/activateEmblem", handler.ActivateEmblem)
		api.POST("/emblem/fetchEmblem", handler.FetchEmblem)
		api.POST("/emblem/fetchEmblemById", handler.FetchEmblemById)

		api.POST("/gameSettings/updatePushNotificationSettings", handler.UpdatePushNotificationSettings)

		api.POST("/friend/fetchFriendList", handler.FetchFriendList)

		api.POST("/lesson/executeLesson", handler.ExecuteLesson)
		api.POST("/lesson/resultLesson", handler.ResultLesson)
		api.POST("/lesson/saveDeck", handler.SaveDeckLesson)
		api.POST("/lesson/skillEditResult", handler.SkillEditResult)
		api.POST("/lesson/changeDeckNameLessonDeck", handler.ChangeDeckNameLessonDeck)

		api.POST("/liveDeck/fetchLiveDeckSelect", handler.FetchLiveDeckSelect)
		api.POST("/liveDeck/changeDeckNameLiveDeck", handler.ChangeDeckNameLiveDeck)
		api.POST("/liveDeck/saveDeckAll", handler.SaveDeckAll)
		api.POST("/liveDeck/saveDeck", handler.SaveDeck)
		api.POST("/liveDeck/saveSuit", handler.SaveSuit)
		api.POST("/livePartners/fetch", live.FetchLivePartners)
		api.POST("/livePartners/setLivePartner", handler.SetLivePartner)

		api.POST("/live/fetchLiveMusicSelect", live.FetchLiveMusicSelect)
		api.POST("/live/start", live.LiveStart)
		api.POST("/live/finish", live.LiveFinish)
		api.POST("/live/skip", live.LiveSkip)
		api.POST("/live/updatePlayList", live.LiveUpdatePlayList)
		api.POST("/live/finishTutorial", live.LiveFinish) // this works

		api.POST("/liveMv/saveDeck", handler.LiveMvSaveDeck)
		api.POST("/liveMv/start", handler.LiveMvStart)

		api.POST("/tutorial/corePlayableEnd", handler.CorePlayableEnd)
		api.POST("/tutorial/phaseEnd", handler.PhaseEnd)
		api.POST("/tutorial/tutorialSkip", handler.TutorialSkip)
		api.POST("/tutorial/timingAdjusterEnd", handler.TimingAdjusterEnd)

		api.POST("/login/login", handler.Login)
		api.POST("/login/startup", handler.Startup)

		api.POST("/loginBonus/readLoginBonus", handler.ReadLoginBonus)

		api.POST("/mission/clearMissionBadge", handler.ClearMissionBadge)
		api.POST("/mission/fetchMission", handler.FetchMission)

		api.POST("/navi/saveUserNaviVoice", handler.SaveUserNaviVoice)
		api.POST("/navi/tapLovePoint", handler.TapLovePoint)

		api.POST("/notice/fetchNotice", handler.FetchNotice)

		api.POST("/present/fetch", handler.FetchPresent)

		// /referenceBook/ all done
		api.POST("/referenceBook/saveReferenceBook", handler.SaveReferenceBook)
		api.POST("/ruleDescription/saveRuleDescription", handler.SaveRuleDescription)

		api.POST("/sif2DataLink/dataLink", handler.Sif2DataLink)
		api.POST("/still/fetch", handler.FetchStill)

		// /story/ all done
		api.POST("/story/finishStoryLinkage", handler.FinishStoryLinkage)
		api.POST("/story/finishUserStoryMain", handler.FinishStoryMain)
		api.POST("/story/saveBrowseStoryMainDigestMovie", handler.SaveBrowseStoryMainDigestMovie)

		// /storyEventHistory/ all done
		api.POST("/storyEventHistory/unlockStory", handler.UnlockStory)
		api.POST("/storyEventHistory/finishStory", handler.FinishStory)

		api.POST("/subscription/fetchSubscriptionPass", handler.FetchSubscriptionPass)

		api.POST("/terms/agreement", handler.Agreement)

		api.POST("/infoTrigger/read", handler.TriggerRead)
		api.POST("/infoTrigger/readCardGradeUp", handler.TriggerReadCardGradeUp)
		api.POST("/infoTrigger/readMemberLoveLevelUp", handler.TriggerReadMemberLoveLevelUp)
		api.POST("/infoTrigger/readMemberGuildSupportItemExpired", handler.TriggerReadMemberGuildSupportItemExpired)

		api.POST("/trainingTree/fetchTrainingTree", handler.FetchTrainingTree)
		api.POST("/trainingTree/levelUpCard", handler.LevelUpCard)
		api.POST("/trainingTree/gradeUpCard", handler.GradeUpCard)
		api.POST("/trainingTree/activateTrainingTreeCell", handler.ActivateTrainingTreeCell)

		api.POST("/userAccountDeletion/checkUserAccountDeleted", handler.CheckUserAccountDeleted)

		api.POST("/userProfile/fetchProfile", handler.FetchProfile)
		api.POST("/userProfile/setProfile", handler.SetProfile)
		api.POST("/userProfile/setProfileBirthday", handler.SetProfileBirthday)
		api.POST("/userProfile/setRecommendCard", handler.SetRecommendCard)
		api.POST("/userProfile/setScoreLive", handler.SetScoreOrComboLive)
		api.POST("/userProfile/setCommboLive", handler.SetScoreOrComboLive) // setCommboLive is a typo in the client?

		api.POST("/gdpr/updateConsentState", handler.UpdateConsentState)
		api.POST("/member/openMemberLovePanel", handler.OpenMemberLovePanel)

		api.POST("/gacha/fetchGachaMenu", handler.FetchGachaMenu)
		api.POST("/gacha/draw", handler.GachaDraw)

		api.POST("/accessory/updateIsLock", handler.AccessoryUpdateIsLock)
		api.POST("/accessory/updateIsNew", handler.AccessoryUpdateIsNew)
		api.POST("/accessory/melt", handler.AccessoryMelt)
		api.POST("/accessory/powerUp", handler.AccessoryPowerUp)
		api.POST("/accessory/rarityUp", handler.AccessoryRarityUp)
		api.POST("/accessory/allUnequip", handler.AccessoryAllUnequip)

		api.POST("/trade/fetchTrade", handler.FetchTrade)
		api.POST("/trade/executeTrade", handler.ExecuteTrade)
		api.POST("/trade/executeMultiTrade", handler.ExecuteMultiTrade)

		api.POST("/takeOver/checkTakeOver", handler.CheckTakeOver)
		api.POST("/takeOver/setTakeOver", handler.SetTakeOver)
		api.POST("/takeOver/updatePassWord", handler.UpdatePassWord)
		api.POST("/dataLink/fetchDataLinks", handler.FetchDataLinks)

		api.POST("/shop/fetchShopTop", handler.FetchShopTop)
		api.POST("/shop/fetchShopPack", handler.FetchShopPack)
		api.POST("/shop/fetchShopSnsCoin", handler.FetchShopSnsCoin)
		api.POST("/shop/fetchShopSubscription", handler.FetchShopSubscription)

		api.POST("/loveRanking/fetch", handler.LoveRankingFetch)

		api.POST("/memberGuild/fetchMemberGuildTop", handler.FetchMemberGuildTop)
		api.POST("/memberGuild/fetchMemberGuildSelect", handler.FetchMemberGuildSelect)
		api.POST("/memberGuild/cheerMemberGuild", handler.CheerMemberGuild)
		api.POST("/memberGuild/joinMemberGuild", handler.JoinMemberGuild)
		api.POST("/memberGuild/fetchMemberGuildRanking", handler.FetchMemberGuildRanking)
		api.POST("/memberGuild/fetchMemberGuildRankingYear", handler.FetchMemberGuildRanking)

		api.POST("/dailyTheater/fetchDailyTheater", handler.FetchDailyTheater)
		api.POST("/dailyTheater/setLike", handler.DailyTheaterSetLike)
		api.POST("/dailyTheaterArchive/fetchDailyTheaterArchive", handler.FetchDailyTheaterArchive)

		api.POST("/unlockScene/saveUnlockedScene", handler.SaveUnlockedScene)
		api.POST("/sceneTips/saveSceneTipsType", handler.SaveSceneTipsType)

		// DLP stuff, can be accessed by editing the database to have a DLP session going on
		api.POST("/tower/fetchTowerSelect", handler.FetchTowerSelect)
		api.POST("/tower/fetchTowerTop", handler.FetchTowerTop)
		api.POST("/tower/clearedTowerFloor", handler.ClearedTowerFloor) // dlp story
		api.POST("/tower/recoveryTowerCardUsed", handler.RecoveryTowerCardUsed)
		api.POST("/tower/recoveryTowerCardUsedAll", handler.RecoveryTowerCardUsedAll)
		api.POST("/towerRanking/fetchTowerRanking", handler.FetchTowerRanking)

		// TODO:
		// /schoolIdolFestivalIdReward/fetch: this can be accessed by unlocking a button, cool to see this again especially if you kept your data
		// /voltageRanking/getVoltageRanking: this is the voltage ranking button
	}

	{
		webapi := r.Group("/webui", webui.Common)
		r.Static("/webui", "webui")
		// the web ui cover for functionality that can't be done by the client or is currently missing
		webapi.POST("/birthday", webui.Birthday)
		webapi.POST("/accessory", webui.Accessory)
		webapi.POST("/import_account", webui.ImportAccount)
		webapi.POST("/export_account", webui.ExportAccount)
		webapi.POST("/reset_story_main", webui.ResetProgress)
		webapi.POST("/reset_story_side", webui.ResetProgress)
		webapi.POST("/reset_story_member", webui.ResetProgress)
		webapi.POST("/reset_story_linkage", webui.ResetProgress)
		webapi.POST("/reset_story_event", webui.ResetProgress)
		webapi.POST("/reset_dlp", webui.ResetProgress)
	}
}
