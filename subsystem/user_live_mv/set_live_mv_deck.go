package user_live_mv

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_member"
	"elichika/userdata"

	"reflect"
)

func SetLiveMvDeck(session *userdata.Session, liveMasterId, liveMvDeckType int32,
	memberMasterIdByPos generic.Dictionary[int32, int32],
	suitMasterIdByPos generic.Dictionary[int32, generic.Nullable[int32]],
	viewStatusByPos generic.Dictionary[int32, int32]) {
	userLiveMvDeck := client.UserLiveMvDeck{
		LiveMasterId: liveMasterId,
	}

	for pos, memberMasterId := range memberMasterIdByPos.Map {
		reflect.ValueOf(&userLiveMvDeck).Elem().Field(int(pos)).Set(reflect.ValueOf(generic.NewNullable(*memberMasterId)))
	}
	for pos, suitMasterId := range suitMasterIdByPos.Map {
		reflect.ValueOf(&userLiveMvDeck).Elem().Field(12 + int(pos)).Set(reflect.ValueOf(*suitMasterId))
	}
	for pos, viewStatus := range viewStatusByPos.Map {
		memberId := memberMasterIdByPos.GetOnly(pos)
		if *memberId == enum.MemberMasterIdRina {
			RinaChan := user_member.GetMember(session, enum.MemberMasterIdRina)
			RinaChan.ViewStatus = *viewStatus
			user_member.UpdateMember(session, RinaChan)
		}
	}

	if liveMvDeckType == enum.LiveMvDeckTypeOriginal {
		session.UserModel.UserLiveMvDeckById.Set(liveMasterId, userLiveMvDeck)
	} else {
		session.UserModel.UserLiveMvDeckCustomById.Set(liveMasterId, userLiveMvDeck)
	}

}
