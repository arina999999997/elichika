package user_beginner_challenge

import (
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// return to whether show the button, and the completed list
func GetUserBeginnerChallengeBootstrap(session *userdata.Session) (bool, generic.List[int32]) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTutorialEnd {
		return false, generic.List[int32]{}
	}
	completedIds := FetchChallengeBeginner(session).CompletedIds
	return completedIds.Size() != len(session.Gamedata.BeginnerChallenge), completedIds
}
