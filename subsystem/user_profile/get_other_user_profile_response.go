package user_profile

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfileResponse(session *userdata.Session, otherUserId int32) response.UserProfileResponse {
	resp := response.UserProfileResponse{
		ProfileInfo: GetOtherUserProfileInfomation(session, otherUserId),
		GuestInfo:   GetOtherUserProfileGuestConfig(session, otherUserId),
		PlayInfo:    GetOtherUserProfilePlayHistory(session, otherUserId),
	}

	cards := []client.UserCard{}
	err := session.Db.Table("u_card").Where("user_id = ?", otherUserId).Find(&cards)
	utils.CheckErr(err)

	ownedCardCount := map[int32]int32{}
	allTrainingActivatedCardCount := map[int32]int32{}
	for _, card := range cards {
		memberId := session.Gamedata.Card[card.CardMasterId].Member.Id
		ownedCardCount[memberId]++
		if card.IsAllTrainingActivated {
			allTrainingActivatedCardCount[memberId]++
		}
	}

	members := []client.UserMember{}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserId).OrderBy("member_master_id").Find(&members)
	utils.CheckErr(err)

	for _, member := range members {
		resp.MemberInfo.UserMembers.Append(client.ProfileUserMember{
			MemberMasterId:                member.MemberMasterId,
			LoveLevel:                     member.LoveLevel,
			LovePointLimit:                member.LovePointLimit,
			OwnedCardCount:                ownedCardCount[member.MemberMasterId],
			AllTrainingActivatedCardCount: allTrainingActivatedCardCount[member.MemberMasterId],
		})
	}

	utils.CheckErr(err)
	return resp
}
