package locale

import (
	"elichika/config"
	"elichika/dictionary"
	"elichika/gamedata"
	"elichika/serverdata"
	"elichika/utils"

	"xorm.io/xorm"
)

type Locale struct {
	// Loaded bool
	Path          string
	Language      string
	StartupKey    []byte
	MasterVersion string
	Gamedata      *gamedata.Gamedata
	Dictionary    *dictionary.Dictionary
}

func (locale *Locale) Load() {
	var err error
	MasterdataEngine, err := xorm.NewEngine("sqlite", locale.Path+"masterdata.db")
	utils.CheckErr(err)
	MasterdataEngine.SetMaxOpenConns(50)
	MasterdataEngine.SetMaxIdleConns(10)
	locale.Dictionary = new(dictionary.Dictionary)
	locale.Dictionary.Init(locale.Path, locale.Language)
	locale.Gamedata = new(gamedata.Gamedata)
	locale.Gamedata.Init(locale.Language, MasterdataEngine, serverdata.Engine, locale.Dictionary)
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
	addLocale(config.GlMasterdataPath, "en", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "zh", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.GlMasterdataPath, "ko", config.MasterVersionGl, config.GlStartupKey)
	addLocale(config.JpMasterdataPath, "ja", config.MasterVersionJp, config.JpStartupKey)
}
