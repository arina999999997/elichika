package user_story_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_live_difficulty"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"
)

// return true if this is first clear
// insert the story if necessary
func FinishStoryMember(session *userdata.Session, storyMemberMasterId int32) {
	storyMemberMaster := session.Gamedata.StoryMember[storyMemberMasterId]

	userStoryMember := client.UserStoryMember{}
	exist, err := session.Db.Table("u_story_member").Where("user_id = ? AND story_member_master_id = ?",
		session.UserId, storyMemberMasterId).Get(&userStoryMember)
	utils.CheckErr(err)
	if !exist {
		userStoryMember = client.UserStoryMember{
			StoryMemberMasterId: storyMemberMasterId,
			IsNew:               true,
			AcquiredAt:          session.Time.Unix(),
		}
		session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)

	}
	if userStoryMember.IsNew {
		userStoryMember.IsNew = false
		if storyMemberMaster.Reward != nil {
			user_present.AddPresent(session, client.PresentItem{
				Content:          *storyMemberMaster.Reward,
				PresentRouteType: enum.PresentRouteTypeStoryMember,
				PresentRouteId:   generic.NewNullable(storyMemberMasterId),
			})
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeStoryMemberReward,
				ParamInt:        generic.NewNullable(storyMemberMasterId),
			})
		}
	}
	// always try to unlock the live
	if storyMemberMaster.UnlockLiveId != nil {
		user_live_difficulty.UnlockLive(session, *storyMemberMaster.UnlockLiveId)
	}
	session.UserModel.UserStoryMemberById.Set(storyMemberMasterId, userStoryMember)
}
