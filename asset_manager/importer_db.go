package asset_manager

import (
	"elichika/log"
)

// TODO(pack): maybe import pack here, for now it's lazily imported later on

func ImportGenericAssetFromDatabase[Specifier AssetTypeSpecifier](assetDatabase *AssetDatabase, locale string) {
	var specifier Specifier
	slice, good := assetDatabase.GenericAssetsByTableName[specifier.TableName()].([]GenericAsset[Specifier])
	if !good {
		return
	}
	for i := range slice {
		slice[i].Locale = locale
		RegisterRawAsset(&slice[i])
	}
}

func ImportAssetFromDatabase(path, locale string) {
	log.Printf("%s %s\n", path, locale)
	db := NewAssetDatabaseForImporting(path)
	for _, pack := range db.Packs {
		RegisterPack(pack)
	}
	for i := range db.Sounds {
		db.Sounds[i].Locale = locale
		RegisterRawAsset(db.Sounds[i])
	}
	for i := range db.Movies {
		db.Movies[i].Locale = locale
		RegisterRawAsset(db.Movies[i])
	}
	ImportGenericAssetFromDatabase[AdvScriptSpecifier](db, locale)
	ImportGenericAssetFromDatabase[BackgroundSpecifier](db, locale)
	ImportGenericAssetFromDatabase[GachaPerformanceSpecifier](db, locale)
	ImportGenericAssetFromDatabase[Live2dSdModelSpecifier](db, locale)
	ImportGenericAssetFromDatabase[LivePropSkeletonSpecifier](db, locale)
	ImportGenericAssetFromDatabase[LiveTimelineSpecifier](db, locale)
	ImportGenericAssetFromDatabase[MemberFacialSpecifier](db, locale)
	ImportGenericAssetFromDatabase[MemberFacialAnimationSpecifier](db, locale)
	ImportGenericAssetFromDatabase[MemberModelSpecifier](db, locale)
	ImportGenericAssetFromDatabase[MemberSdModelSpecifier](db, locale)
	ImportGenericAssetFromDatabase[NaviMotionSpecifier](db, locale)
	ImportGenericAssetFromDatabase[NaviTimelineSpecifier](db, locale)
	ImportGenericAssetFromDatabase[ShaderSpecifier](db, locale)
	ImportGenericAssetFromDatabase[SkillEffectSpecifier](db, locale)
	ImportGenericAssetFromDatabase[SkillTimelineSpecifier](db, locale)
	ImportGenericAssetFromDatabase[SkillWipeSpecifier](db, locale)
	ImportGenericAssetFromDatabase[StageSpecifier](db, locale)
	ImportGenericAssetFromDatabase[StageEffectSpecifier](db, locale)
	ImportGenericAssetFromDatabase[TextureSpecifier](db, locale)
	FreeAssetDatabase(db)
}
