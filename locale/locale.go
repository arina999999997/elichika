package locale

import (
	"elichika/assetdata"
	"elichika/config"
	"elichika/db"
	"elichika/dictionary"
	"elichika/gamedata"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"
	"time"

	"xorm.io/xorm"
)

// create one engine for each potential file being read
// each locale is free to create and store its own session
var engines = map[string]*xorm.Engine{}

func getEngine(path string) *xorm.Engine {
	engine, exist := engines[path]
	if exist {
		return engine
	}
	engine, err := xorm.NewEngine("sqlite", path)
	utils.CheckErr(err)
	engines[path] = engine
	return engine
}

type Locale struct {
	Path          string
	Language      string
	StartupKey    []byte
	MasterVersion string
	Gamedata      *gamedata.Gamedata
	Dictionary    *dictionary.Dictionary
}

func (locale *Locale) LoadGamedata(syncChannel chan struct{}) {
	locale.Dictionary = new(dictionary.Dictionary)
	locale.Dictionary.Init(locale.Path, locale.Language)
	locale.Gamedata = new(gamedata.Gamedata)
	masterdataDb, err := db.NewDatabase(locale.Path + "masterdata.db")
	utils.CheckErr(err)
	locale.Gamedata.Init(locale.Language, masterdataDb, serverdata.Database, locale.Dictionary, syncChannel)

}

func (locale *Locale) LoadAsset() {
	AssetdataEngine := getEngine(fmt.Sprintf("%s/asset_a_%s.db", locale.Path, locale.Language))
	AssetdataEngine.SetMaxOpenConns(50)
	AssetdataEngine.SetMaxIdleConns(10)
	assetdata.Init(locale.Language, AssetdataEngine)

	AssetdataEngine = getEngine(fmt.Sprintf("%s/asset_i_%s.db", locale.Path, locale.Language))
	AssetdataEngine.SetMaxOpenConns(50)
	AssetdataEngine.SetMaxIdleConns(10)
	assetdata.Init(locale.Language, AssetdataEngine)
}

var (
	Locales map[string](*Locale)
)

func addLocale(path, language, masterVersion, startUpKey string) {
	locale := Locale{
		Path:          path,
		Language:      language,
		MasterVersion: masterVersion,
		StartupKey:    []byte(startUpKey),
	}
	Locales[language] = &locale
}

func init() {
	start := time.Now()
	gamedata.GenerateLoadOrder()
	Locales = make(map[string](*Locale))
	syncChannel := make(chan struct{})
	addLocale(config.JpMasterdataPath, "ja", config.MasterVersionJp, config.JpStartupKey)
	addLocale(config.GlMasterdataPath, "en", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "zh", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "ko", config.MasterVersionGl, config.GlStartupKey)

	for _, locale := range Locales {
		go locale.LoadGamedata(syncChannel)
	}
	// asset write to the same space so needed to be load manually
	for _, locale := range Locales {
		locale.LoadAsset()
	}
	for i := len(Locales); i > 0; i-- {
		<-syncChannel
	}
	finish := time.Now()
	fmt.Println("Finished loading databases in: ", finish.Sub(start))
	for language, locale := range Locales {
		gamedata.GamedataByLocale[language] = locale.Gamedata
		// because the order of has map is random, this instance is guaranteed to not
		// be a specific version, so don't depend on it
		gamedata.Instance = locale.Gamedata
	}
}
