package locale

import (
	"elichika/assetdata"
	"elichika/config"
	"elichika/dictionary"
	"elichika/gamedata"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"

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

func (locale *Locale) Load() {
	MasterdataEngine := getEngine(locale.Path + "masterdata.db")
	MasterdataEngine.SetMaxOpenConns(50)
	MasterdataEngine.SetMaxIdleConns(10)
	locale.Dictionary = new(dictionary.Dictionary)
	locale.Dictionary.Init(locale.Path, locale.Language)
	locale.Gamedata = new(gamedata.Gamedata)
	locale.Gamedata.Init(locale.Language, MasterdataEngine, serverdata.Engine, locale.Dictionary)
	// asset data is shared among all locale, but it is initiated multiple times, each time adding more data
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
	locale.Load()
	Locales[language] = &locale
}

func init() {
	Locales = make(map[string](*Locale))
	addLocale(config.JpMasterdataPath, "ja", config.MasterVersionJp, config.JpStartupKey)
	addLocale(config.GlMasterdataPath, "en", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "zh", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "ko", config.MasterVersionGl, config.GlStartupKey)
}
