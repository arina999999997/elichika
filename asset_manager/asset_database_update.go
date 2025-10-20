package asset_manager

import (
	"elichika/utils"

	"encoding/hex"
	"fmt"
	"math/rand"
)

// return the sql for updating assets, with a comment stating if an asset from different locale was used
func GetUpdateSQLForAssets(assetDatabasePath string, assetPaths []string, preferedLocale string, backupLocales []string) string {
	// check first
	for _, assetPath := range assetPaths {
		_, exists := Assets[assetPath]
		if !exists {
			panic("Asset path is not loaded and cannot be used " + assetPath)
		}
	}
	// for now we will generate the result in the order of passed in assetPaths
	// maybe we can add some sorting helper in the future to get a specific order
	packUpdates := ""
	assetPathUpdates := ""
	assetDatabase := NewAssetDatabaseForUpdating(assetDatabasePath)
	// gather the necessary pack and insert them
	packCheck := map[string]bool{}

	for _, assetPath := range assetPaths {
		rawAsset := Assets[assetPath].GetRawAssetByLocale(preferedLocale)
		if rawAsset == nil {
			rawAsset = Assets[assetPath].GetRawAssetByPreference(backupLocales)
			if rawAsset == nil {
				panic("Asset path %s doesn't have an associated valid raw asset")
			} else {
				assetPathUpdates += fmt.Sprintf("-- no valid asset for locale %s, defaulted to locale %s\n", preferedLocale, rawAsset.GetLocale())
			}
		}
		assetPathUpdates += rawAsset.GetUpdateSQL(assetDatabase)
		// gather the pack
		pack := rawAsset.GetPack()
		if pack == nil {
			panic("asset doesn't have pack: " + assetPath)
		}
		if !packCheck[pack.PackName] {
			packCheck[pack.PackName] = true
			packUpdates += pack.GetUpdateSQL(assetDatabase)
		}
		secondaryPack := rawAsset.GetAwbPack()
		if (secondaryPack != nil) && (!packCheck[secondaryPack.PackName]) {
			packCheck[secondaryPack.PackName] = true
			packUpdates += secondaryPack.GetUpdateSQL(assetDatabase)
		}
	}
	// finally merge the asset update
	if packUpdates != "" {
		// recount the amount of pack associated with main key
		// for now just use random version, it will work
		// but we can probably setup some sort of hashing, but that require reading more data from the package

		version := make([]byte, 20)
		_, err := rand.Read(version)
		utils.CheckErr(err)
		packUpdates += fmt.Sprintf("UPDATE m_asset_package SET version = \"%s\", pack_num = (SELECT COUNT(*) FROM m_asset_package_mapping WHERE package_key=\"main\") WHERE package_key = \"main\";\n", hex.EncodeToString(version))
	}
	return packUpdates + assetPathUpdates
}
