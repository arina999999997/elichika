package user_login_bonus

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"fmt"
	"math/rand"
	"time"
)

func birthdayLoginBonusHandler(mode string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
	if loginBonus.LoginBonusType != enum.LoginBonusTypeBirthday {
		panic("wrong handler used")
	}
	year, month, day := session.Time.Date()
	mmdd := int32(month)*100 + int32(day)
	list, exists := session.Gamedata.MemberByBirthday[mmdd]
	fmt.Println(mmdd, list, exists)
	if !exists { // no one with this birthday
		return
	}
	userLoginBonus := getUserLoginBonus(session, loginBonus.LoginBonusId)
	lastUnlocked := time.Date(year, month, day, 0, 0, 0, 0, session.Time.Location())
	if userLoginBonus.LastReceivedAt >= lastUnlocked.Unix() { // already got it
		return
	}
	userLoginBonus.LastReceivedAt = session.Time.Unix()

	for _, member := range list {
		// the present is as follow:
		// - 50 gems
		// - 2 memento
		// - 50 memorial
		// - additional 25 memorial for channel member
		naviLoginBonus := loginBonus.NaviLoginBonus()
		naviLoginBonus.LoginBonusRewards.Append(
			client.LoginBonusRewards{
				Day:          1,
				Status:       enum.LoginBonusReceiveStatusReceiving,
				ContentGrade: generic.NewNullable(enum.LoginBonusContentGradeRare),
				LoginBonusContents: generic.Array[client.Content]{
					Slice: []client.Content{
						client.Content{
							ContentType:   enum.ContentTypeTrainingMaterial,
							ContentId:     int32(18000 + member.Id),
							ContentAmount: 2,
						},
						client.Content{
							ContentType:   enum.ContentTypeTrainingMaterial,
							ContentId:     int32(8000 + member.Id),
							ContentAmount: 50,
						},
						item.StarGem.Amount(50),
					},
				},
			},
		)
		if session.UserStatus.MemberGuildMemberMasterId.HasValue &&
			session.UserStatus.MemberGuildMemberMasterId.Value == member.Id {
			naviLoginBonus.LoginBonusRewards.Slice[0].LoginBonusContents.Slice[1].ContentAmount += 25
		}

		for _, content := range naviLoginBonus.LoginBonusRewards.Slice[0].LoginBonusContents.Slice {
			user_present.AddPresent(session, client.PresentItem{
				Content:          content,
				PresentRouteType: enum.PresentRouteTypeLoginBonus,
				PresentRouteId:   generic.NewNullable(int32(1000003)),
				// TODO(localization): This is not localized to the correct language, also we don't even know if this is the correct string
				ParamServer: generic.NewNullable(client.LocalizedText{
					DotUnderText: fmt.Sprint(member.Name, " Birthday Login Bonus"),
				}),
			})
		}

		// choose the background and the costume
		memberLoginBonusBirthday := member.MemberLoginBonusBirthdays[0]
		switch mode {
		case "random":
			memberLoginBonusBirthday = member.MemberLoginBonusBirthdays[rand.Intn(len(member.MemberLoginBonusBirthdays))]
		case "latest":
		default:
			panic("not supported")
		}
		target.BirthdayMember.Append(client.LoginBonusBirthDayMember{
			MemberMasterId: generic.NewNullable(member.Id),
			SuitMasterId:   generic.NewNullable(memberLoginBonusBirthday.SuitMasterId),
		})
		naviLoginBonus.BackgroundId = memberLoginBonusBirthday.Id
		target.BirthdayLoginBonuses.Append(naviLoginBonus)
	}
	updateUserLoginBonus(session, userLoginBonus)
}
