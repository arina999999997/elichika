package user_member_guild

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// Check the previous member guild for potential reward
// - we only check once and mark the check
// - then we create a trigger for this
// - the trigger is only removed once reward are delivered
// - return the trigger and whether it's valid
func CheckPreviousMemberGuildReward(session *userdata.Session) (int64, bool) {
	previousMemberGuildId := GetCurrentMemberGuildId(session) - 1
	if previousMemberGuildId <= 0 {
		return 0, false
	}

	lastReceived := database.UserMemberGuildRankingRewardReceived{}
	exist, err := session.Db.Table("u_member_guild_ranking_reward_received").
		Where("user_id = ?", session.UserId).Get(&lastReceived)
	utils.CheckErr(err)
	if lastReceived.MemberGuildId < previousMemberGuildId { // there might be reward
		lastReceived.MemberGuildId = previousMemberGuildId
		// always update so we don't check again, then we check the actual reward
		if exist {
			_, err := session.Db.Table("u_member_guild_ranking_reward_received").Where("user_id = ?", session.UserId).
				AllCols().Update(&lastReceived)
			utils.CheckErr(err)
		} else {
			userdata.GenericDatabaseInsert(session, "u_member_guild_ranking_reward_received", lastReceived)
		}

		// requirement for reward is that the point gained > 0
		userMemberGuild := GetUserMemberGuild(session, previousMemberGuildId)
		if userMemberGuild.TotalPoint > 0 {
			_, resultAt := GetMemberGuildStartAndEnd(session, previousMemberGuildId)
			limitAt := resultAt + session.Gamedata.MemberGuildPeriod.OneCycleSecs
			triggerId := session.NextUniqueId()
			user_info_trigger.AddTriggerBasic(session, client.UserInfoTriggerBasic{
				TriggerId:       triggerId,
				InfoTriggerType: enum.InfoTriggerTypeMemberGuildRankingShowResult,
				LimitAt:         generic.NewNullable(limitAt),
			})
			return triggerId, true
		}

	} else { // the trigger might still be here
		resultTriggers := []client.UserInfoTriggerBasic{}
		err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ?",
			session.UserId, enum.InfoTriggerTypeMemberGuildRankingShowResult).Find(&resultTriggers)
		utils.CheckErr(err)
		for _, trigger := range resultTriggers {
			if trigger.LimitAt.Value >= session.Time.Unix() {
				return trigger.TriggerId, true
			}
		}
	}
	return 0, false
}
