// this package represents the gamedata
// i.e. settings, parameters of the games that are shared between all users
// the data are stored in both masterdata.db and serverdata.db
// some data are loaded only from one of the 2, but some data need boths
//
// no handling function should interact with the master/server data db and only interact with this package
// this is done both to reduce the time necessary to look into db, as well as to simplify accessing data to the most
// relevant id, instead of having to access multiple tables or use magic id system
// for example, everything related to a single card / accessory will use that card / accessory master id as id
// everything related to all card / accessory of a rarity will use that rarity as id
// relation is defined by masterdata.db or serverdata.db
// i.e. some setting might be the same across all accessory of a rarity, but as long as it's store separately in the db,
// it's stored separately here

// id priority:
// - if there is an exclusive id, use it unless it doesn't matter.
// - otherwise the id priority go from general to specific
// - for example, the member id is used more than the member love ids, so the member id would be the outer access

// referenced object
//   - if an object reference another object, then the data structure should store a reference to the object referenced by that id
//   - the object should keep a reference to the id itself
//   - for example, m_live_difficulty reference m_live through live_id
//     then m_live_difficulty contain LiveId *int, Live *Live
//   - LiveId will be loaded into at first, then it should point to Live.LiveId after we setup Live
package gamedata

import (
	"elichika/client"
	"elichika/db"
	"elichika/dictionary"
	"elichika/generic/drop"

	"reflect"
)

type loadFunc = func(*Gamedata)

var (
	funcs       map[uintptr]loadFunc
	prequisites map[uintptr][]uintptr
	loadOrder   []loadFunc
)

func addLoadFunc(f loadFunc) {
	if funcs == nil {
		funcs = make(map[uintptr]loadFunc)
		prequisites = make(map[uintptr][]uintptr)
	}
	funcs[reflect.ValueOf(f).Pointer()] = f
}

func addPrequisite(function, prequisite loadFunc) {
	addLoadFunc(function)
	addLoadFunc(prequisite)
	prequisites[reflect.ValueOf(function).Pointer()] = append(prequisites[reflect.ValueOf(function).Pointer()],
		reflect.ValueOf(prequisite).Pointer())
}

func generateLoadOrder(fid uintptr) {
	_, exist := funcs[fid]
	if !exist {
		return // done
	}
	for _, prequisite := range prequisites[fid] {
		generateLoadOrder(prequisite)
	}

	loadOrder = append(loadOrder, funcs[fid])
	delete(funcs, fid)
}

type Gamedata struct {
	Language     string
	ServerdataDb *db.DatabaseSync
	MasterdataDb *db.DatabaseSync
	Dictionary   *dictionary.Dictionary

	Accessory              map[int32]*Accessory
	AccessoryRarity        map[int32]*AccessoryRarity
	AccessoryRarityUpGroup map[int32]*AccessoryRarityUpGroup
	AccessoryMeltGroup     map[int32]*AccessoryMeltGroup
	AccessoryLevelUpItem   map[int32]*AccessoryLevelUpItem

	ActivityPointRecoveryPrice []ActivityPointRecoveryPrice

	BeginnerChallenge                         map[int32]*BeginnerChallenge
	BeginnerChallengeCell                     map[int32]*BeginnerChallengeCell
	BeginnerChallengeCellByClearConditionType map[int32][]*BeginnerChallengeCell

	NgWord *NgWordNode

	ConstantInt []int32

	Emblem map[int32]*Emblem

	EventActive         EventActive
	EventAvailable      EventAvailable
	EventMarathon       map[int32]*EventMarathon
	EventMarathonReward map[int32][]*client.Content
	EventStory          map[int32]*EventStory

	Member                          map[int32]*Member
	MemberGroup                     map[int32]*MemberGroup
	MemberLoveLevelLovePoint        []int32
	MemberLoveLevelCount            int32
	MemberLovePanel                 map[int32]*MemberLovePanel
	MemberLovePanelCell             map[int32]*MemberLovePanelCell
	MemberLovePanelLevelAtLoveLevel []int32
	MemberByBirthday                map[int32]([]*Member)

	MemberGuildConstant             MemberGuildConstant
	MemberGuildPeriod               MemberGuildPeriod
	MemberGuildCheerReward          map[int32]*drop.DropList[client.Content]
	MemberGuildRankingRewardInside  map[int32]*MemberGuildRankingReward
	MemberGuildRankingRewardOutside map[int32]*MemberGuildRankingReward
	MemberGuildPointClearReward     map[int32]*MemberGuildPointClearReward
	MemberGuildChallengeLive        MemberGuildChallengeLive

	Mission                     map[int32]*Mission
	MissionByClearConditionType map[int32][]*Mission
	MissionByTerm               map[int32][]*Mission
	MissionByTriggerType        map[int32][]*Mission

	NaviVoice map[int32]*NaviVoice

	CustomBackground map[int32]*CustomBackground

	Live                 map[int32]*Live
	LiveParty            LiveParty
	LiveDaily            map[int32]*LiveDaily
	LiveMemberMapping    map[int32]LiveMemberMapping
	LiveDifficulty       map[int32]*LiveDifficulty
	LiveDropGroup        map[int32]*LiveDropGroup
	LiveDropContentGroup map[int32]*drop.WeightedDropList[client.Content]

	LessonMenu map[int32]*LessonMenu

	TrainingTreeCellItemSet map[int32]*TrainingTreeCellItemSet
	TrainingTreeDesign      map[int32]*TrainingTreeDesign
	TrainingTreeMapping     map[int32]*TrainingTreeMapping
	TrainingTree            map[int32]*TrainingTree

	Card           map[int32]*Card
	CardLevel      map[int32]*CardLevel
	CardRarity     map[int32]*CardRarity
	CardByMemberId map[int32][]*Card

	ContentRarity         *ContentRarity
	ContentType           map[int32]*ContentType
	Content               map[int32]map[int32]*Content // content type then content id, not guaranteed to be valid for things like Card/Suit/...
	ContentsByContentType map[int32][]*client.Content

	Suit map[int32]*Suit

	StoryMember       map[int32]*StoryMember
	StoryMainChapter  map[int32]*StoryMainChapter
	StoryEventHistory map[int32]*StoryEventHistory

	Gacha              map[int32]*Gacha
	GachaList          []*Gacha
	GachaDraw          map[int32]*GachaDraw
	GachaDrawGuarantee map[int32]*GachaDrawGuarantee
	GachaGroup         map[int32]*GachaGroup
	GachaGuarantee     map[int32]*GachaGuarantee

	Tower map[int32]*Tower

	Trade        map[int32]*Trade   // map from TradeId to Trade
	TradesByType [3][]*client.Trade // map from trade type to array of Trade
	TradeProduct map[int32]*client.TradeProduct

	LoginBonus map[int32]*LoginBonus

	UserRank    map[int32]*UserRank
	UserRankMax int32

	DailyTheater          map[int32]*DailyTheater
	LastestDailyTheaterId int32
}

// for convinience of access, these are provided
// they are only valid after the locale system write into them itself
// allow for session-less access to gamedata, at the cost of being
// locale non-specific
var Instance *Gamedata = nil
var GamedataByLocale = map[string]*Gamedata{}

func GenerateLoadOrder() {
	for len(funcs) > 0 {
		var fid uintptr
		for key := range funcs {
			fid = key
			break
		}
		generateLoadOrder(fid)
	}
}

func (gamedata *Gamedata) Init(language string, masterdata *db.DatabaseSync, serverdata *db.DatabaseSync, dictionary *dictionary.Dictionary, syncChannel chan struct{}) {
	gamedata.Language = language
	// gamedata.MasterdataDb = masterdata.NewSession()
	gamedata.MasterdataDb = masterdata
	gamedata.ServerdataDb = serverdata
	gamedata.Dictionary = dictionary

	for _, loadFunc := range loadOrder {
		loadFunc(gamedata)
	}
	syncChannel <- struct{}{}
}
