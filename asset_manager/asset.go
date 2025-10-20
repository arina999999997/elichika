package asset_manager

// asset class, accessed by asset path alone
// this refer to an abstract asset and can be of any type (i.e. texture, sound, ...)
import (
	"elichika/log"

	"fmt"
)

type Asset struct {
	AssetPath        string
	RawAssetByLocale map[string]RawAsset
}

func (asset *Asset) RegisterRawAsset(rawAsset RawAsset) {
	locale := rawAsset.GetLocale()
	existingRawAsset, exists := asset.RawAssetByLocale[locale]
	if exists {
		// log.Printf("Potentially conflicting asset: \"%s\", current data \"%s\", new data: \"%s\" \n",
		// asset.AssetPath, (*existingRawAsset).Info(), rawAsset.Info())
		if existingRawAsset.IsAvailable() {
			log.Printf("There's already a valid raw asset for asset path \"%s\" and locale \"%s\": \"%s\", ignoring new data: \"%s\" \n",
				asset.AssetPath, locale, existingRawAsset.Info(), rawAsset.Info())
			return
		}
		if rawAsset.IsAvailable() {
			log.Printf("Updated to new valid raw asset for asset path \"%s\" and locale \"%s\": from \"%s\" to \"%s\" \n",
				asset.AssetPath, locale, existingRawAsset.Info(), rawAsset.Info())
			asset.RawAssetByLocale[locale] = rawAsset
			return
		}
		log.Printf("Duplicated invalid asset: \"%s\" and locale \"%s\": from \"%s\" to \"%s\" \n",
			asset.AssetPath, locale, existingRawAsset.Info(), rawAsset.Info())
	} else {
		asset.RawAssetByLocale[locale] = rawAsset
	}
}

// note that these functions return a RawAsset if and only if it is available (has local file)
// even if the reference to that RawAsset is known from the database, it doesn't return if it isn't valid

func (asset Asset) GetRawAssetByLocale(locale string) RawAsset {
	rawAsset := asset.RawAssetByLocale[locale]
	if rawAsset == nil {
		return nil
	}
	if rawAsset.IsAvailable() {
		return rawAsset
	} else {
		return nil
	}
}

func (asset Asset) GetRawAssetAny() RawAsset {
	for _, rawAsset := range asset.RawAssetByLocale {
		if rawAsset.IsAvailable() {
			return rawAsset
		}
	}
	return nil
}

func (asset Asset) GetRawAssetByPreference(preferences []string) RawAsset {
	for _, locale := range preferences {
		rawAsset, exists := asset.RawAssetByLocale[locale]
		if exists && rawAsset.IsAvailable() {
			return rawAsset
		}
	}
	return nil
}

func (asset Asset) Info() string {
	res := fmt.Sprint("AssetPath: ", asset.AssetPath, " Locale Data: ")
	for locale, rawAsset := range asset.RawAssetByLocale {
		res += fmt.Sprintf("(Locale: %s, Info: %s)", locale, rawAsset.Info())
	}
	return res
}
