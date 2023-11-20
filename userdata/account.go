package userdata

import (
	"elichika/dictionary"
	"elichika/gamedata"
	"elichika/model"
	"elichika/utils"

	// "encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
	// "github.com/tidwall/sjson"
)

// return the userID if it is not given
func CreateNewAccount(ctx *gin.Context, userID int, passWord string) int {
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	dictionary := ctx.MustGet("dictionary").(*dictionary.Dictionary)
	{
		db := Engine.NewSession()
		defer db.Close()
		err := db.Begin()
		utils.CheckErr(err)
		isRandomID := (userID == -1)
		if isRandomID {
			userID = rand.Intn(1000000000)
		}
		status := model.UserStatus{
			UserID:                                  userID,
			PassWord:                                passWord,
			LastLoginAt:                             time.Now().Unix(),
			Rank:                                    1,
			Exp:                                     0,
			RecommendCardMasterID:                   100011001, // Honoka
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
			LatestLiveDeckID:                          1,
			MainLessonDeckID:                          1,
			FavoriteMemberID:                          1,
			LastLiveDifficultyID:                      10001101,
			LpMagnification:                           1,
			EmblemID:                                  10500521, // new player
			DeviceToken:                               "",
			TutorialPhase:                             99,
			TutorialEndAt:                             1622217482,
			LoginDays:                                 1221,
			NaviTapCount:                              0,
			NaviTapRecoverAt:                          1688137200,
			IsAutoMode:                                false,
			MaxScoreLiveDifficultyMasterID:            10001101,
			LiveMaxScore:                              0,
			MaxComboLiveDifficultyMasterID:            10001101,
			LiveMaxCombo:                              0,
			LessonResumeStatus:                        1,
			AccessoryBoxAdditional:                    400,
			TermsOfUseVersion:                         2,
			BootstrapSifidCheckAt:                     1692471111782,
			GdprVersion:                               4,
			MemberGuildMemberMasterID:                 1,
			MemberGuildLastUpdatedAt:                  1659485328,
			Cash:                                      0,
		}
		status.Name.DotUnderText = "Newcomer"
		status.Nickname.DotUnderText = "Newcomer"
		status.Message.DotUnderText = "Hello!"
		// insert into the db
		_, err = db.Table("u_info").AllCols().Insert(&status)
		if (err != nil) && (isRandomID) { // reroll once for random userID
			userID = rand.Intn(1000000000)
			status.UserID = userID
			_, err = db.Table("u_info").AllCols().Insert(&status)
		}
		utils.CheckErr(err)
		db.Commit()
	}
	session := GetSession(ctx, userID)
	defer session.Close()

	{ // members, initial cards
		members := []model.UserMember{}
		cards := []model.UserCard{}

		for _, member := range gamedata.Member {
			members = append(members, model.UserMember{
				UserID:                   userID,
				MemberMasterID:           member.ID,
				CustomBackgroundMasterID: member.MemberInit.CustomBackgroundMID,
				SuitMasterID:             member.MemberInit.SuitMasterID,
				LovePoint:                member.MemberInit.LovePoint,
				LoveLevel:                member.MemberInit.LoveLevel,
				LovePointLimit:           member.MemberInit.LovePointLimit,
				ViewStatus:               1,
				IsNew:                    false,
				OwnedCardCount:           1,
				AllTrainingCardCount:     0,
			})
			cards = append(cards, model.UserCard{
				UserID:                     userID,
				CardMasterID:               member.MemberInit.SuitMasterID,
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
				AdditionalPassiveSkill1ID:  0,
				AdditionalPassiveSkill2ID:  0,
				AdditionalPassiveSkill3ID:  0,
				AdditionalPassiveSkill4ID:  0,
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
					UserID:       userID,
					SuitMasterID: suit.ID,
					IsNew:        false,
				})
			}
		}
		session.InsertUserSuits(suits)
	}
	{ // show formation
		liveDecks := []model.UserLiveDeck{}
		liveParties := []model.UserLiveParty{}
		for i := 1; i <= 20; i++ {
			cid := [10]int{}
			// this order isn't actually correct to the official server
			for j := 1; j <= 9; j++ {
				cid[j] = gamedata.Member[j+100*((i-1)%3)].MemberInit.SuitMasterID
			}
			liveDeck := model.UserLiveDeck{
				UserID:         userID,
				UserLiveDeckID: i,
				CardMasterID1:  cid[1],
				CardMasterID2:  cid[2],
				CardMasterID3:  cid[3],
				CardMasterID4:  cid[4],
				CardMasterID5:  cid[5],
				CardMasterID6:  cid[6],
				CardMasterID7:  cid[7],
				CardMasterID8:  cid[8],
				CardMasterID9:  cid[9],
				SuitMasterID1:  cid[1],
				SuitMasterID2:  cid[2],
				SuitMasterID3:  cid[3],
				SuitMasterID4:  cid[4],
				SuitMasterID5:  cid[5],
				SuitMasterID6:  cid[6],
				SuitMasterID7:  cid[7],
				SuitMasterID8:  cid[8],
				SuitMasterID9:  cid[9],
			}
			liveDeck.Name.DotUnderText = fmt.Sprintf(dictionary.Resolve("k.m_sorter_deck_live")+" %d", i)
			liveDecks = append(liveDecks, liveDeck)
			for j := 1; j <= 3; j++ {
				liveParty := model.UserLiveParty{
					UserID:         userID,
					PartyID:        i*100 + j,
					UserLiveDeckID: i,
					CardMasterID1:  cid[(j-1)*3+1],
					CardMasterID2:  cid[(j-1)*3+2],
					CardMasterID3:  cid[(j-1)*3+3],
				}
				liveParty.IconMasterID, liveParty.Name.DotUnderText = gamedata.
					GetLivePartyInfoByCardMasterIDs(liveParty.CardMasterID1, liveParty.CardMasterID2, liveParty.CardMasterID3)
				liveParties = append(liveParties, liveParty)
			}
		}
		session.InsertLiveDecks(liveDecks)
		session.InsertLiveParties(liveParties)
	}
	{ // lesson deck

		lessonDecks := []model.UserLessonDeck{}
		for i := 1; i <= 20; i++ {
			lessonDecks = append(lessonDecks, model.UserLessonDeck{
				UserID:           userID,
				UserLessonDeckID: i,
				Name:             fmt.Sprintf(dictionary.Resolve("k.m_sorter_deck_lesson")+" %d", i),
			})
		}
		session.InsertLessonDecks(lessonDecks)
	}
	session.Finalize("{}", "user_model")

	return userID
}
