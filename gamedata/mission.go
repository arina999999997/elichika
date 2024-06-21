package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/enum"
	"elichika/utils"

	"fmt"
	"sort"
	"time"

	"xorm.io/xorm"
)

type Mission struct {
	// from m_mission
	Id   int32 `xorm:"pk 'id'"`
	Term int32 `xorm:"'term'" enum:"MissionTerm"`
	// Title string `xorm:"'title'"`
	// Description string `xorm:"'description'"`
	TriggerType       int32 `xorm:"'trigger_type'" enum:"MissionTrigger"`
	TriggerCondition1 int32 `xorm:"'trigger_condition_1'"`
	// TriggerCondition2 *int32 `xorm:"'trigger_condition_2'"` // always null
	StartAt int64 `xorm:"'start_at'"`
	EndAt   int64 `xorm:"'end_at'"` // can be null, but the loader will set this to 2^63 - 1 if it is null
	// SceneTransitionLink int32 `xorm:"'scene_transition_link'"`
	// SceneTransitionParam *int32 `xorm:"'scene_transition_param'"`
	PickupType *int32 `xorm:"'pickup_type'" enum:"MissionPickupType"`
	// DisplayOrder int32 `xorm:"'display_order'"`
	MissionClearConditionType   int32  `xorm:"'mission_clear_condition_type'" enum:"MissionClearConditionType"`
	MissionClearConditionCount  int32  `xorm:"'mission_clear_condition_count'"`
	MissionClearConditionParam1 *int32 `xorm:"'mission_clear_condition_param1'"`
	MissionClearConditionParam2 *int32 `xorm:"'mission_clear_condition_param2'"` // always null
	CompleteMissionNum          *int32 `xorm:"'complete_mission_num'"`           // only for beginner missions
	// HasContent bool `xorm:"'has_content'"` // always false

	TriggerMissions []*Mission `xorm:"-"`

	// from m_mission_reward
	Rewards []client.Content `xorm:"-"`
}

func (m *Mission) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := masterdata_db.Table("m_mission_reward").Where("mission_id = ?", m.Id).OrderBy("display_order").Find(&m.Rewards)
	utils.CheckErr(err)

	if m.EndAt == 0 {
		m.EndAt = (1 << 63) - 1
	}
}

func loadMission(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Mission")
	gamedata.Mission = make(map[int32]*Mission)
	err := masterdata_db.Table("m_mission").Find(&gamedata.Mission)
	utils.CheckErr(err)
	gamedata.MissionByClearConditionType = make(map[int32][]*Mission)
	gamedata.MissionByTerm = make(map[int32][]*Mission)
	gamedata.MissionByTriggerType = make(map[int32][]*Mission)
	for _, mission := range gamedata.Mission {
		mission.populate(gamedata, masterdata_db, serverdata_db, dictionary)

		gamedata.MissionByClearConditionType[mission.MissionClearConditionType] =
			append(gamedata.MissionByClearConditionType[mission.MissionClearConditionType], mission)
		gamedata.MissionByTerm[mission.Term] =
			append(gamedata.MissionByTerm[mission.Term], mission)
		gamedata.MissionByTriggerType[mission.TriggerType] =
			append(gamedata.MissionByTriggerType[mission.TriggerType], mission)

		if mission.TriggerType == enum.MissionTriggerClearMission {
			gamedata.Mission[mission.TriggerCondition1].TriggerMissions = append(gamedata.Mission[mission.TriggerCondition1].TriggerMissions, mission)
			if mission.MissionClearConditionType != gamedata.Mission[mission.TriggerCondition1].MissionClearConditionType {
				if mission.CompleteMissionNum == nil {
					panic(fmt.Sprint("different clear contition type from parent mission ", mission.Id, mission.TriggerCondition1))
				}
			}
		} else if mission.TriggerType != enum.MissionTriggerGameStart {
			panic("unsupported trigger type")
		}
	}
	for _, list := range gamedata.MissionByClearConditionType {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Id < list[j].Id
		})
	}
	for _, list := range gamedata.MissionByTerm {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Id < list[j].Id
		})
	}
	for _, list := range gamedata.MissionByTriggerType {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Id < list[j].Id
		})
	}
	// the complete all daily / weekly mission don't have correct goal, we have to count it manually
	timeStamp := time.Now().Unix()
	for _, mission := range gamedata.MissionByTerm[enum.MissionTermDaily] {
		if mission.EndAt < timeStamp {
			continue
		}
		if mission.MissionClearConditionType == enum.MissionClearConditionTypeCompleteDaily {
			mission.MissionClearConditionCount = 0
			for _, m := range gamedata.MissionByTerm[enum.MissionTermDaily] {
				if (m.EndAt < timeStamp) || (m == mission) {
					continue
				}
				if (m.PickupType == nil) != (mission.PickupType == nil) {
					continue
				}
				if (mission.PickupType == nil) || (*mission.PickupType == *m.PickupType) {
					mission.MissionClearConditionCount++
				}
			}
		}
	}
	for _, mission := range gamedata.MissionByTerm[enum.MissionTermWeekly] {
		if mission.EndAt < timeStamp {
			continue
		}
		if mission.MissionClearConditionType == enum.MissionClearConditionTypeCompleteWeekly {
			mission.MissionClearConditionCount = 0
			for _, m := range gamedata.MissionByTerm[enum.MissionTermWeekly] {
				if (m.EndAt < timeStamp) || (m == mission) {
					continue
				}
				if (m.PickupType == nil) != (mission.PickupType == nil) {
					continue
				}
				if (mission.PickupType == nil) || (*mission.PickupType == *m.PickupType) {
					mission.MissionClearConditionCount++
				}
			}
		}
	}
}

func init() {
	addLoadFunc(loadMission)
}
