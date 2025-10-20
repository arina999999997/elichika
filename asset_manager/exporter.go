package asset_manager

// export the info about asset to database files:
// - asset_manager_cached_asset_en.db
// - asset_manager_cached_asset_ja.db
// - asset_manager_cached_asset_ko.db
// - asset_manager_cached_asset_zh.db
// - asset_manager_cached_asset_un.db
// the format is similar to the asset_a_en.db files, except without some relevant pack

import (
	"elichika/log"
	"elichika/utils"

	"xorm.io/xorm"
)

func CachedDBName(locale string) string {
	return "asset_manager_cached_asset_" + locale + ".db"
}

func BackupCachedDatabase(locale string) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		log.Println("Error backing up for locale: ", locale, err)
	}()
	utils.CopyFile(CachedDBName(locale), CachedDBName(locale)+".backup")
}

func SetupCachedDatabaseTable(engine *xorm.Engine, tableName string, structure interface{}) {
	exist, err := engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)
	if exist {
		err = engine.DropTables(tableName)
		utils.CheckErr(err)
	}
	err = engine.Table(tableName).CreateTable(structure)
	utils.CheckErr(err)
}

func SetupCachedDatabaseGenericAsset[Specifier AssetTypeSpecifier](engine *xorm.Engine) {
	var specifier Specifier
	SetupCachedDatabaseTable(engine, specifier.TableName(), GenericAsset[Specifier]{})
}

func SetupCachedDatabase(engine *xorm.Engine) {
	SetupCachedDatabaseTable(engine, "m_asset_sound", Sound{})
	SetupCachedDatabaseTable(engine, "m_movie", Movie{})
	SetupCachedDatabaseTable(engine, "m_asset_package_mapping", Pack{})
	SetupCachedDatabaseGenericAsset[AdvScriptSpecifier](engine)
	SetupCachedDatabaseGenericAsset[BackgroundSpecifier](engine)
	SetupCachedDatabaseGenericAsset[GachaPerformanceSpecifier](engine)
	SetupCachedDatabaseGenericAsset[Live2dSdModelSpecifier](engine)
	SetupCachedDatabaseGenericAsset[LivePropSkeletonSpecifier](engine)
	SetupCachedDatabaseGenericAsset[LiveTimelineSpecifier](engine)
	SetupCachedDatabaseGenericAsset[MemberFacialSpecifier](engine)
	SetupCachedDatabaseGenericAsset[MemberFacialAnimationSpecifier](engine)
	SetupCachedDatabaseGenericAsset[MemberModelSpecifier](engine)
	SetupCachedDatabaseGenericAsset[MemberSdModelSpecifier](engine)
	SetupCachedDatabaseGenericAsset[NaviMotionSpecifier](engine)
	SetupCachedDatabaseGenericAsset[NaviTimelineSpecifier](engine)
	SetupCachedDatabaseGenericAsset[ShaderSpecifier](engine)
	SetupCachedDatabaseGenericAsset[SkillEffectSpecifier](engine)
	SetupCachedDatabaseGenericAsset[SkillTimelineSpecifier](engine)
	SetupCachedDatabaseGenericAsset[SkillWipeSpecifier](engine)
	SetupCachedDatabaseGenericAsset[StageSpecifier](engine)
	SetupCachedDatabaseGenericAsset[StageEffectSpecifier](engine)
	SetupCachedDatabaseGenericAsset[TextureSpecifier](engine)
}

func ExportCachedData() {
	// backup the database if there's any
	locales := []string{"en", "ja", "ko", "zh", "un"}
	for _, locale := range locales {
		BackupCachedDatabase(locale)
	}
	AssetDatabaseByLocale := map[string]*AssetDatabase{}
	for _, locale := range locales {
		log.Printf("Setting up target database for locale: %s\n", locale)
		engine, err := xorm.NewEngine("sqlite", CachedDBName(locale))
		utils.CheckErr(err)
		SetupCachedDatabase(engine)
		err = engine.Close()
		utils.CheckErr(err)
		AssetDatabaseByLocale[locale] = NewAssetDatabaseForExporting(CachedDBName(locale))
	}
	for _, asset := range Assets {
		for locale, rawAsset := range asset.RawAssetByLocale {
			log.Println(rawAsset.Info())
			rawAsset.InsertIntoAssetDatabase(AssetDatabaseByLocale[locale])
		}
	}
	for _, assetDatabase := range AssetDatabaseByLocale {
		assetDatabase.ExportToDb()
	}
}
