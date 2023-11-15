// this package represents the gamedata
// i.e. settings, parameters of the games that are shared between all users
// the data are stored in both masterdata.db and serverdata.db
// some data are loaded only from one of the 2, but some data need boths
//
// eventually no handling function should interact with the db and only interact with this package
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
//     then m_live_difficulty contain LiveID *int, Live *Live
//   - LiveID will be loaded into at first, then it should point to Live.LiveID after we setup Live
package gamedata

import (
	"elichika/dictionary"
	"elichika/model"

	"reflect"
	// "fmt"

	"xorm.io/xorm"
)

type loadFunc = func(*Gamedata, *xorm.Session, *xorm.Session, *dictionary.Dictionary)

var (
	funcs       map[uintptr]loadFunc
	prequisites map[uintptr][]uintptr
	loadOrder   []loadFunc
)

func addLoadFunc(name loadFunc) {
	if funcs == nil {
		funcs = make(map[uintptr]loadFunc)
		prequisites = make(map[uintptr][]uintptr)
	}
	funcs[reflect.ValueOf(name).Pointer()] = name
}

func addPrequisite(function, prequisite loadFunc) {
	addLoadFunc(function)
	addLoadFunc(prequisite)
	prequisites[reflect.ValueOf(function).Pointer()] = append(prequisites[reflect.ValueOf(function).Pointer()],
		reflect.ValueOf(prequisite).Pointer())
}

func generateLoadOrder(fid uintptr) {
	_, exists := funcs[fid]
	if !exists {
		return // done
	}
	for _, prequisite := range prequisites[fid] {
		generateLoadOrder(prequisite)
	}

	loadOrder = append(loadOrder, funcs[fid])
	delete(funcs, fid)
}

type Gamedata struct {
	Accessory              map[int]*Accessory
	AccessoryRarity        map[int]*AccessoryRarity
	AccessoryRarityUpGroup map[int]*AccessoryRarityUpGroup
	AccessoryMeltGroup     map[int]*AccessoryMeltGroup
	AccessoryLevelUpItem   map[int]*AccessoryLevelUpItem

	Member                          map[int]*Member
	MemberLoveLevelLovePoint        []int
	MemberLoveLevelCount            int
	MemberLovePanel                 map[int]*MemberLovePanel
	MemberLovePanelCell             map[int]*MemberLovePanelCell
	MemberLovePanelLevelAtLoveLevel []int

	Live              map[int]*Live
	LiveParty         LiveParty
	LiveDaily         map[int]*LiveDaily
	LiveMemberMapping map[int]LiveMemberMapping
	LiveDifficulty    map[int]*LiveDifficulty

	TrainingTreeCellItemSet map[int]*TrainingTreeCellItemSet
	TrainingTreeMapping     map[int]*TrainingTreeMapping
	TrainingTree            map[int]*TrainingTree

	Card      map[int]*Card
	CardLevel map[int]*CardLevel

	Suit map[int]*Suit

	Gacha          map[int]*Gacha
	GachaList      []*Gacha
	GachaDraw      map[int]*GachaDraw
	GachaGroup     map[int]*GachaGroup
	GachaGuarantee map[int]*GachaGuarantee

	Trade        map[int]*model.Trade // map from TradeID to Trade
	TradesByType [3][]*model.Trade    // map from trade type to array of Trade
	TradeProduct map[int]*model.TradeProduct
}

func (gamedata *Gamedata) Init(masterdata *xorm.Engine, serverdata *xorm.Engine, dictionary *dictionary.Dictionary) {
	masterdata_session := masterdata.NewSession()
	serverdata_session := serverdata.NewSession()
	defer masterdata_session.Close()
	defer serverdata_session.Close()

	for len(funcs) > 0 {
		var fid uintptr
		for key := range funcs {
			fid = key
			break
		}
		generateLoadOrder(fid)
	}
	for _, loadFunc := range loadOrder {
		loadFunc(gamedata, masterdata_session, serverdata_session, dictionary)
	}
}
