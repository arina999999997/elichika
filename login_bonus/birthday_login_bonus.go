package login_bonus

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/item"
	"elichika/userdata"

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
	if !exists { // no one with this birthday
		return
	}
	userLoginBonus := session.GetUserLoginBonus(loginBonus.LoginBonusId)
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
		naviLoginBonus.LoginBonusRewards = append(naviLoginBonus.LoginBonusRewards,
			client.LoginBonusRewards{
				Day:          1,
				Status:       enum.LoginBonusReceiveStatusReceiving,
				ContentGrade: generic.NewNullable(enum.LoginBonusContentGradeRare),
				LoginBonusContents: []client.Content{
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
		)
		if *session.UserStatus.MemberGuildMemberMasterId == int32(member.Id) {
			naviLoginBonus.LoginBonusRewards[0].LoginBonusContents[1].ContentAmount += 25
		}

		for _, content := range naviLoginBonus.LoginBonusRewards[0].LoginBonusContents {
			// TODO(present_box): This correctly has to go to the present box, but we just do it here
			session.AddResource(content)
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
		target.BirthdayMember = append(target.BirthdayMember, client.LoginBonusBirthDayMember{
			MemberMasterId: generic.NewNullable(int32(member.Id)),
			SuitMasterId:   generic.NewNullable(int32(memberLoginBonusBirthday.SuitMasterId)),
		})
		naviLoginBonus.BackgroundId = int32(memberLoginBonusBirthday.Id)
		target.BirthdayLoginBonuses = append(target.BirthdayLoginBonuses, naviLoginBonus)
	}
	session.UpdateUserLoginBonus(userLoginBonus)
}
