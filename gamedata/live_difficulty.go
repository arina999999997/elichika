package gamedata

import (
	"elichika/client"
	"elichika/config"
	"elichika/dictionary"
	"elichika/enum"
	"elichika/generic"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"sort"

	"xorm.io/xorm"
)

type LiveDifficulty struct {
	// from m_live_difficulty
	LiveDifficultyId int32  `xorm:"pk 'live_difficulty_id'"`
	LiveId           *int32 `xorm:"'live_id'"`
	Live             *Live  `xorm:"-"`
	// Live3DAssetMasterId *int
	LiveDifficultyType int32 `xorm:"'live_difficulty_type'" enum:""`
	UnlockPattern      int32 `xorm:"'unlock_pattern'" enum:""`
	// DefaultAttribute int32
	TargetVoltage int32 `xorm:"'target_voltage'"`
	NoteEmitMsec  int32 `xorm:"'note_emit_msec'"`
	ConsumedLP    int32 `xorm:"'consumed_lp'"`
	RewardUserExp int32 `xorm:"'reward_user_exp'"`
	// JudgeId int32
	NoteDropGroupId *int32 `xorm:"'note_drop_group_id'"`

	// NoteDropGroup *NoteDropGroup `xorm:"-"`
	DropChooseCount    int32  `xorm:"'drop_choose_count'"`
	RateDropRate       int32  `xorm:"'rare_drop_rate'"`
	DropContentGroupId *int32 `xorm:"'drop_content_group_id'"`
	// DropContentGroup *DropContentGroup `xorm:"-"`
	RareDropContentGroupId *int32 `xorm:"'rare_drop_content_group_id'"`
	// RareDropContentGroup *RareDropContentGroup `xorm:"-"`
	AdditionalDropContentGroupId *int32 `xorm:"'additional_drop_content_group_id'"`
	// AdditionalDropContentGroup *AdditionalDropContentGroup `xorm:"-"`
	// ?????
	BottomTechnique              int32 `xorm:"'bottom_technique'"`
	AdditionalDropDecayTechnique int32 `xorm:"'additional_drop_decay_technique'"`

	RewardBaseLovePoint int32 `xorm:"'reward_base_love_point'"`
	EvaluationSScore    int32 `xorm:"'evaluation_s_score'"`
	EvaluationAScore    int32 `xorm:"'evaluation_a_score'"`
	EvaluationBScore    int32 `xorm:"'evaluation_b_score'"`
	EvaluationCScore    int32 `xorm:"'evaluation_c_score'"`
	// UpdatedAt int `xorm:"'updated_at'"`
	LoseAtDeath bool `xorm:"'lose_at_death'"`
	// AutoplayRequirementId *int `xorm:"'autoplay_requirement_id'"`
	SkipMasterId *int32 `xorm:"'skip_master_id'"`
	// StaminaVoltageGroupId int
	// ComboVoltageGroupId int
	// DifficultyConstMasterId int
	IsCountTarget bool `xorm:"'is_count_target'"`
	// InsufficentRate int

	// from m_live_difficulty_mission
	Missions []LiveDifficultyMission `xorm:"-"`

	// lazily constructed?
	LiveStage       *client.LiveStage `xorm:"-"`
	SimpleLiveStage *SimpleLiveStage  `xorm:"-"`

	// from m_live_difficulty_gimmick
	LiveDifficultyGimmick *LiveDifficultyGimmick `xorm:"-"`

	// from m_live_difficulty_note_gimmick
	LiveDifficultyNoteGimmicks []LiveDifficultyNoteGimmick `xorm:"-"`
}

func (ld *LiveDifficulty) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	ld.Live = gamedata.Live[*ld.LiveId]
	// 2-way links
	ld.Live.LiveDifficulties = append(ld.Live.LiveDifficulties, ld)
	ld.LiveId = &gamedata.Live[*ld.LiveId].LiveId
	err := masterdata_db.Table("m_live_difficulty_mission").Where("live_difficulty_master_id = ?", ld.LiveDifficultyId).
		OrderBy("position").Find(&gamedata.LiveDifficulty[ld.LiveDifficultyId].Missions)
	utils.CheckErr(err)
	// if ld.LiveDifficultyId == 9999 || ld.LiveDifficultyId/10 == 6000000  {
	// 	return
	// }

	ld.LiveDifficultyGimmick = new(LiveDifficultyGimmick)
	exists, err := masterdata_db.Table("m_live_difficulty_gimmick").Where("live_difficulty_master_id = ?", ld.LiveDifficultyId).
		Get(ld.LiveDifficultyGimmick)
	utils.CheckErr(err)

	if !exists {
		// doesn't exist for a small set of things that shouldn't matter
		// panic(fmt.Sprint("gimmick doesn't exist for: ", ld.LiveDifficultyId))
		ld.LiveDifficultyGimmick = nil
		// fmt.Println("gimmick doesn't exist for: ", ld.LiveDifficultyId)
	}

	err = masterdata_db.Table("m_live_difficulty_note_gimmick").Where("live_difficulty_id = ?", ld.LiveDifficultyId).
		Find(&ld.LiveDifficultyNoteGimmicks)
	utils.CheckErr(err)
	for i := range ld.LiveDifficultyNoteGimmicks {
		ld.LiveDifficultyNoteGimmicks[i].populate()
	}
}

func (ld *LiveDifficulty) loadSimpleLiveStage(gamedata *Gamedata) {
	if ld.SimpleLiveStage != nil {
		return // already loaded
	}
	// fmt.Println("Loading for", ld.LiveDifficultyId)
	liveNotes := utils.ReadAllText(fmt.Sprintf("assets/simple_stages/%d.json", ld.LiveDifficultyId))
	if (liveNotes == "") || (ld.UnlockPattern == enum.LiveUnlockPatternTowerOnly) {

		// song doesn't exist, use rule to find the original map
		if ld.UnlockPattern != enum.LiveUnlockPatternTowerOnly {
			// only accept event songs, SBL, or DLP
			return
		}
		originalLiveId := ld.Live.LiveId%10000 + 10000
		for _, other := range gamedata.Live[originalLiveId].LiveDifficulties {
			if (other.NoteEmitMsec == ld.NoteEmitMsec) && (other.LiveDifficultyType == ld.LiveDifficultyType) {
				other.loadSimpleLiveStage(gamedata)
				if other.SimpleLiveStage != nil {
					ld.SimpleLiveStage = other.SimpleLiveStage
					break
				}
			}
		}
		if ld.SimpleLiveStage == nil {
			for _, other := range gamedata.Live[originalLiveId].LiveDifficulties {
				if other.NoteEmitMsec == ld.NoteEmitMsec {
					other.loadSimpleLiveStage(gamedata)
					if other.SimpleLiveStage != nil {
						ld.SimpleLiveStage = other.SimpleLiveStage
						break
					}
				}
			}
		}
	} else {
		err := json.Unmarshal([]byte(liveNotes), &ld.SimpleLiveStage)
		utils.CheckErr(err)
	}
	if ld.SimpleLiveStage == nil {
		panic(fmt.Sprint("Error finding live stage for: ", ld.LiveDifficultyId))
	}
	if ld.SimpleLiveStage.Original != nil {
		_, exists := gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original]
		if !exists {
			fmt.Println("Warning: original live referenced but do not exist in database: ",
				*ld.SimpleLiveStage.Original, ". Attemping to just load the json.")
			gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original] = new(LiveDifficulty)
			gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original].LiveDifficultyId = *ld.SimpleLiveStage.Original
			gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original].LiveDifficultyType = ld.LiveDifficultyType
		}
		gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original].loadSimpleLiveStage(gamedata)
		ld.SimpleLiveStage = gamedata.LiveDifficulty[*ld.SimpleLiveStage.Original].SimpleLiveStage
	}
	if ld.SimpleLiveStage == nil {
		panic(fmt.Sprint("Error finding original live stage for: ", ld.LiveDifficultyId))
	}
}

func (ld *LiveDifficulty) ConstructLiveStage(gamedata *Gamedata) {
	if ld.LiveStage != nil { // generated
		return
	}

	if !config.GenerateStageFromScratch { // load generated stage, it must exists
		text := utils.ReadAllText(fmt.Sprintf("assets/stages/%d.json", ld.LiveDifficultyId))
		if text == "" {
			panic(fmt.Sprintf("Stage %d doesn't exists in assets/stages", ld.LiveDifficultyId))
		}
		ld.LiveStage = new(client.LiveStage)
		err := json.Unmarshal([]byte(text), &ld.LiveStage)
		if err != nil {
			panic(fmt.Sprintf("Failed to load stage %d: wrong format", ld.LiveDifficultyId))
		}
		return
	}

	ld.loadSimpleLiveStage(gamedata)
	if ld.SimpleLiveStage == nil {
		if ld.UnlockPattern != enum.LiveUnlockPatternTowerOnly {
			return
		}
		panic(fmt.Sprint("Failed to load simple live stage for: ", ld.LiveDifficultyId))
	}

	// make the object and set relevant stuff
	ld.LiveStage = new(client.LiveStage)
	ld.LiveStage.LiveDifficultyId = ld.LiveDifficultyId

	ld.LiveStage.LiveNotes.Slice = append(ld.LiveStage.LiveNotes.Slice, ld.SimpleLiveStage.LiveNotes...)
	for i := range ld.LiveStage.LiveNotes.Slice {
		ld.LiveStage.LiveNotes.Slice[i].Id = int32(i + 1)
		ld.LiveStage.LiveNotes.Slice[i].AutoJudgeType = enum.JudgeTypeGreat         // can be overwritten at runtime
		ld.LiveStage.LiveNotes.Slice[i].NoteRandomDropColor = enum.NoteDropColorNon // can be overwritten at runtime
	}
	ld.LiveStage.LiveWaveSettings.Slice = append(ld.LiveStage.LiveWaveSettings.Slice, ld.SimpleLiveStage.LiveWaveSettings...)

	// each note store its own gimmick, and the stage store unique note gimmicks in it
	noteGimmickDict := map[int32]bool{}
	for _, noteGimmick := range ld.LiveDifficultyNoteGimmicks {
		ld.LiveStage.LiveNotes.Slice[noteGimmick.NoteId-1].GimmickId = noteGimmick.Id
		if !noteGimmickDict[noteGimmick.Id] {
			noteGimmickDict[noteGimmick.Id] = true
			ld.LiveStage.NoteGimmicks.Append(client.NoteGimmick{
				Id:              noteGimmick.Id,
				NoteGimmickType: noteGimmick.NoteGimmickType,
				EffectMId:       noteGimmick.SkillMasterId,
				IconType:        noteGimmick.NoteGimmickIconType,
			})
		}
	}
	sort.Slice(ld.LiveStage.NoteGimmicks.Slice, func(i, j int) bool {
		return ld.LiveStage.NoteGimmicks.Slice[i].Id < ld.LiveStage.NoteGimmicks.Slice[j].Id
	})
	for i := range ld.LiveStage.NoteGimmicks.Slice {
		ld.LiveStage.NoteGimmicks.Slice[i].UniqId = int32(2001 + i)
	}
	if ld.LiveDifficultyGimmick != nil {
		ld.LiveStage.StageGimmickDict.Set(ld.LiveDifficultyGimmick.TriggerType, generic.Array[client.LiveStageGimmick]{
			Slice: []client.LiveStageGimmick{client.LiveStageGimmick{
				GimmickMasterId:    ld.LiveDifficultyGimmick.Id,
				ConditionMasterId1: ld.LiveDifficultyGimmick.ConditionMasterId1,
				ConditionMasterId2: generic.NewNullable(ld.LiveDifficultyGimmick.ConditionMasterId2),
				SkillMasterId:      ld.LiveDifficultyGimmick.SkillMasterId,
				UniqId:             1001,
			},
			}})

	}

	// save the new map
	{
		output, err := json.Marshal(ld.LiveStage)
		utils.CheckErr(err)
		utils.WriteAllText(fmt.Sprintf("assets/stages/%d.json", ld.LiveDifficultyId), string(output))
	}

	// check against pregenerated map
	// skip checking for coop (SBL), because the database only has constant modifier while the actual
	// data will have some added bonus gimmick
	// not like we use those map right now anyway
	if ld.UnlockPattern == enum.LiveUnlockPatternCoopOnly {
		return
	}
	text := utils.ReadAllText(fmt.Sprintf("assets/full_stages/%d.json", ld.LiveDifficultyId))
	if text == "" {
		// fmt.Println("Newly generated map: ", ld.LiveDifficultyId)
		return
	}
	pregeneratedStage := client.LiveStage{}
	err := json.Unmarshal([]byte(text), &pregeneratedStage)
	utils.CheckErr(err)
	if !pregeneratedStage.IsSame(ld.LiveStage) {
		panic(fmt.Sprint("Difference detected for: ", ld.LiveDifficultyId, "\n", ld.LiveStage, "\n_______________\n", pregeneratedStage))
	}

}

func loadLiveDifficulty(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDifficulty")
	gamedata.LiveDifficulty = make(map[int32]*LiveDifficulty)
	err := masterdata_db.Table("m_live_difficulty").Find(&gamedata.LiveDifficulty)
	utils.CheckErr(err)
	// ordered iteration is important here
	ids := []int32{}
	for id := range gamedata.LiveDifficulty {
		ids = append(ids, id)
	}
	// order by unlock pattern then id
	sort.Slice(ids, func(i, j int) bool {
		if gamedata.LiveDifficulty[ids[i]].UnlockPattern != gamedata.LiveDifficulty[ids[j]].UnlockPattern {
			return gamedata.LiveDifficulty[ids[i]].UnlockPattern < gamedata.LiveDifficulty[ids[j]].UnlockPattern
		}
		return gamedata.LiveDifficulty[ids[i]].LiveDifficultyId < gamedata.LiveDifficulty[ids[j]].LiveDifficultyId
	})
	for _, id := range ids {
		liveDifficulty := gamedata.LiveDifficulty[id]
		liveDifficulty.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}

	if config.GenerateStageFromScratch {
		for _, liveDifficulty := range gamedata.LiveDifficulty {
			liveDifficulty.ConstructLiveStage(gamedata)
		}
	}
}

func init() {
	addLoadFunc(loadLiveDifficulty)
	addPrequisite(loadLiveDifficulty, loadLive)
}
