package userdata

import (
	"elichika/config"
	"elichika/dictionary"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/utils"

	// "encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// return the userId if it is not given
func CreateNewAccount(ctx *gin.Context, userId int, passWord string) int {
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	dictionary := ctx.MustGet("dictionary").(*dictionary.Dictionary)
	{
		db := Engine.NewSession()
		defer db.Close()
		err := db.Begin()
		utils.CheckErr(err)
		isRandomId := (userId == -1)
		if isRandomId {
			userId = rand.Intn(1000000000)
		}
		tutorialPhase := 99
		tutorialEndAt := int(time.Now().Unix())
		if *config.Conf.Tutorial {
			tutorialPhase = enum.TutorialPhaseNameInput
			tutorialEndAt = 0
		}
		status := model.UserStatus{
			UserId:                                  userId,
			PassWord:                                passWord,
			LastLoginAt:                             time.Now().Unix(),
			Rank:                                    1,
			RecommendCardMasterId:                   100011001, // Honoka
			MaxFriendNum:                            99,
			LivePointFullAt:                         time.Now().Unix(),
			LivePointBroken:                         10000,
			LivePointSubscriptionRecoveryDailyCount: 1,
			LivePointSubscriptionRecoveryDailyResetAt: 1688137200,
			ActivityPointCount:                        3,
			ActivityPointResetAt:                      1688137200,
			ActivityPointPaymentRecoveryDailyCount:    10,
			ActivityPointPaymentRecoveryDailyResetAt:  1688137200,
			GameMoney:                                 1000000000,
			CardExp:                                   1000000000,
			FreeSnsCoin:                               50000,
			AppleSnsCoin:                              25000,
			GoogleSnsCoin:                             25000,
			SubscriptionCoin:                          50000,
			BirthDate:                                 202307,
			BirthMonth:                                7,
			BirthDay:                                  1,
			LatestLiveDeckId:                          1,
			MainLessonDeckId:                          1,
			FavoriteMemberId:                          1,
			LastLiveDifficultyId:                      10001101,
			LpMagnification:                           1,
			EmblemId:                                  10500521, // new player
			TutorialPhase:                             tutorialPhase,
			TutorialEndAt:                             tutorialEndAt,
			LoginDays:                                 1221,
			NaviTapRecoverAt:                          1688137200,
			MaxScoreLiveDifficultyMasterId:            10001101,
			MaxComboLiveDifficultyMasterId:            10001101,
			LessonResumeStatus:                        1,
			AccessoryBoxAdditional:                    400,
			TermsOfUseVersion:                         2,
			BootstrapSifidCheckAt:                     1692471111782,
			GdprVersion:                               4,
			MemberGuildMemberMasterId:                 1,
			MemberGuildLastUpdatedAt:                  1659485328,
		}
		status.Name.DotUnderText = "Newcomer"
		status.Nickname.DotUnderText = "Newcomer"
		status.Message.DotUnderText = "Hello!"
		// insert into the db
		_, err = db.Table("u_info").AllCols().Insert(&status)
		if (err != nil) && (isRandomId) { // reroll once for random userId
			userId = rand.Intn(1000000000)
			status.UserId = userId
			_, err = db.Table("u_info").AllCols().Insert(&status)
		}
		utils.CheckErr(err)
		db.Commit()
	}
	session := GetSession(ctx, userId)
	defer session.Close()

	{ // members, initial cards
		members := []model.UserMember{}
		cards := []model.UserCard{}

		for _, member := range gamedata.Member {
			members = append(members, model.UserMember{
				UserId:                   userId,
				MemberMasterId:           member.Id,
				CustomBackgroundMasterId: member.MemberInit.CustomBackgroundMId,
				SuitMasterId:             member.MemberInit.SuitMasterId,
				LovePoint:                member.MemberInit.LovePoint,
				LoveLevel:                member.MemberInit.LoveLevel,
				LovePointLimit:           member.MemberInit.LovePointLimit,
				ViewStatus:               1,
				IsNew:                    false,
				OwnedCardCount:           1,
				AllTrainingCardCount:     0,
			})
			cards = append(cards, model.UserCard{
				UserId:                     userId,
				CardMasterId:               member.MemberInit.SuitMasterId,
				Level:                      1,
				Exp:                        0,
				LovePoint:                  0,
				IsFavorite:                 false,
				IsAwakening:                false,
				IsAwakeningImage:           false,
				IsAllTrainingActivated:     false,
				TrainingActivatedCellCount: 0,
				MaxFreePassiveSkill:        1, // all R
				Grade:                      0,
				TrainingLife:               0,
				TrainingAttack:             0,
				TrainingDexterity:          0,
				ActiveSkillLevel:           1,
				PassiveSkillALevel:         1,
				PassiveSkillBLevel:         1,
				PassiveSkillCLevel:         1,
				AdditionalPassiveSkill1Id:  0,
				AdditionalPassiveSkill2Id:  0,
				AdditionalPassiveSkill3Id:  0,
				AdditionalPassiveSkill4Id:  0,
				AcquiredAt:                 time.Now().Unix(),
				IsNew:                      false,
				LivePartnerCategories:      0,
				LiveJoinCount:              0,
				ActiveSkillPlayCount:       0,
			})
		}
		session.InsertMembers(members)
		session.InsertCards(cards)
	}
	{ // all the costumes that can't be obtained from maxing cards
		suits := []model.UserSuit{}

		for _, suit := range gamedata.Suit {
			if suit.SuitReleaseRoute == 2 {
				suits = append(suits, model.UserSuit{
					UserId:       userId,
					SuitMasterId: suit.Id,
					IsNew:        false,
				})
			}
		}
		session.InsertUserSuits(suits)
	}
	{ // show formation
		liveDecks := []model.UserLiveDeck{}
		liveParties := []model.UserLiveParty{}
		for a := 0; a <= 10000; a += 10000 { // 10000... are DLP formations
			for b := 1; b <= 20; b++ {
				i := a + b
				cid := [10]int{}
				// this order isn't actually correct to the official server
				for j := 1; j <= 9; j++ {
					cid[j] = gamedata.Member[j+100*((i-1)%3)].MemberInit.SuitMasterId
				}
				liveDeck := model.UserLiveDeck{
					UserId:         userId,
					UserLiveDeckId: i,
					CardMasterId1:  cid[1],
					CardMasterId2:  cid[2],
					CardMasterId3:  cid[3],
					CardMasterId4:  cid[4],
					CardMasterId5:  cid[5],
					CardMasterId6:  cid[6],
					CardMasterId7:  cid[7],
					CardMasterId8:  cid[8],
					CardMasterId9:  cid[9],
					SuitMasterId1:  cid[1],
					SuitMasterId2:  cid[2],
					SuitMasterId3:  cid[3],
					SuitMasterId4:  cid[4],
					SuitMasterId5:  cid[5],
					SuitMasterId6:  cid[6],
					SuitMasterId7:  cid[7],
					SuitMasterId8:  cid[8],
					SuitMasterId9:  cid[9],
				}
				liveDeck.Name.DotUnderText = fmt.Sprintf(dictionary.Resolve("k.m_sorter_deck_live")+" %d", i)
				liveDecks = append(liveDecks, liveDeck)
				for j := 1; j <= 3; j++ {
					liveParty := model.UserLiveParty{
						UserId:         userId,
						PartyId:        i*100 + j,
						UserLiveDeckId: i,
						CardMasterId1:  cid[(j-1)*3+1],
						CardMasterId2:  cid[(j-1)*3+2],
						CardMasterId3:  cid[(j-1)*3+3],
					}
					liveParty.IconMasterId, liveParty.Name.DotUnderText = gamedata.
						GetLivePartyInfoByCardMasterIds(liveParty.CardMasterId1, liveParty.CardMasterId2, liveParty.CardMasterId3)
					liveParties = append(liveParties, liveParty)
				}
			}
		}
		session.InsertLiveDecks(liveDecks)
		session.InsertLiveParties(liveParties)
	}
	{ // lesson deck

		lessonDecks := []model.UserLessonDeck{}
		for i := 1; i <= 20; i++ {
			lessonDecks = append(lessonDecks, model.UserLessonDeck{
				UserId:           userId,
				UserLessonDeckId: i,
				Name:             fmt.Sprintf(dictionary.Resolve("k.m_sorter_deck_lesson")+" %d", i),
			})
		}
		session.InsertLessonDecks(lessonDecks)
	}
	session.Finalize("{}", "user_model")

	return userId
}
