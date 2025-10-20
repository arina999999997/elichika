package asset_manager

import (
	"elichika/log"
)

// An asset manager is responsible for holding information about certain assets and performing operations on them.
// To get import into the asset manager, call the RegisterRaw asset and RegisterRawPack manually, or call the premade importer
// - See importer.go for more details on the importer concept
// - See importer_<type>.go for specific importer types
// Assuming the assets are imported, then we can use it to do useful things:
// - Generate SQL update data for the databases:
//   - This require a list of assets, which is user generated.
//   - This can do things like default to a asset from an from a different version / report missing asset
// - Analytic / tool: (TODO for now)
//   - Generate asset package summary:
//     - There might be gaps in the asset files that can contain some missing information
//   - Repacking:
//     - If we already gather up all the relevant assets, we can safely remove old assets to reduce the install size
// Note that due to reason, we won't be using pointer here as the garbage collector doesn't want to free the pointer as much for some reason
// More precisely, things that are modifed are pointer, things that are set and done are value

var (
	RawAssets     = []RawAsset{}
	Assets        = map[string]*Asset{}
	RawAssetCheck = map[string]bool{}

	Packs          = []Pack{}
	PackByPackName = map[string]*Pack{}
)

func RegisterRawAsset(rawAsset RawAsset) {
	if RawAssetCheck[rawAsset.Info()] {
		return
	}
	RawAssetCheck[rawAsset.Info()] = true
	RawAssets = append(RawAssets, rawAsset)
	_, exists := Assets[rawAsset.GetAssetPath()]
	if !exists {
		Assets[rawAsset.GetAssetPath()] = &Asset{
			AssetPath:        rawAsset.GetAssetPath(),
			RawAssetByLocale: map[string]RawAsset{},
		}
	}
	Assets[rawAsset.GetAssetPath()].RegisterRawAsset(rawAsset)
}

func RegisterPack(pack Pack) {
	_, exists := PackByPackName[pack.PackName]
	if exists {
		// TODO(extra): Check if the pack are the same
		return
	}
	Packs = append(Packs, pack)
	PackByPackName[pack.PackName] = &Packs[len(Packs)-1]
}

func GetPack(packName string) *Pack {
	pack, exist := PackByPackName[packName]
	if exist {
		return pack
	}
	log.Printf("WARNING: unknown pack: %s, created dummy pack\n", packName)
	RegisterPack(Pack{
		PackName:       packName,
		FileSize:       27083004,
		MetapackName:   nil,
		MetapackOffset: 0,
	})
	return PackByPackName[packName]
}
