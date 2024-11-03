package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveParty struct {
	// only relevant data for now, full move later on
	partyInfoByRoleIds [5][5][5]struct {
		PartyIcon int32
		PartyName string
	}
}

func (gamedata *Gamedata) GetLivePartyInfoByCardMasterIds(a, b, c int32) (partyIcon int32, partyName string) {
	a = int32(gamedata.Card[a].Role)
	b = int32(gamedata.Card[b].Role)
	c = int32(gamedata.Card[c].Role)
	partyIcon = gamedata.LiveParty.partyInfoByRoleIds[a][b][c].PartyIcon
	partyName = gamedata.LiveParty.partyInfoByRoleIds[a][b][c].PartyName
	return
}

func loadLiveParty(gamedata *Gamedata) {
	fmt.Println("Loading LiveParty")
	type LiveParty struct {
		Id            int    `xorm:"pk 'id'"`
		Role1         int    `xorm:"'role_1'"`
		Role2         int    `xorm:"'role_2'"`
		Role3         int    `xorm:"'role_3'"`
		Name          string `xorm:"'name'"`
		LivePartyIcon int32  `xorm:"'live_party_icon'"`
	}
	liveParties := []LiveParty{}
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_live_party_name").Find(&liveParties)
	})
	utils.CheckErr(err)
	for _, party := range liveParties {
		party.Name = gamedata.Dictionary.Resolve(party.Name)
		r := [3]int{}
		r[0] = party.Role1
		r[1] = party.Role2
		r[2] = party.Role3

		for i := 0; i <= 2; i++ {
			for j := 0; j <= 2; j++ {
				if i == j {
					continue
				}
				k := 3 - i - j
				gamedata.LiveParty.partyInfoByRoleIds[r[i]][r[j]][r[k]].PartyIcon = party.LivePartyIcon
				gamedata.LiveParty.partyInfoByRoleIds[r[i]][r[j]][r[k]].PartyName = party.Name
			}
		}
	}
}

func init() {
	addLoadFunc(loadLiveParty)
}
