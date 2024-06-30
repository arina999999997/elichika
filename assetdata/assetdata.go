package assetdata

// this package manage the asset of the game, or more precisely how to distribute the assets to clients.
// (not to be confused with /assets, which is a git submodule that actually store some of the assets).
// asset follow the rules:
// - each asset is of one of the types that the game use:
//   - for example: texture, sound, stage effect, ...
// - each asset is accessed using an asset path.
// - each asset path will map to a pack:
//   - pack_name: which pack is it
//   - head: where to begin reading the relevant asset in the pack file.
//   - size: the size
//   - key1, key2: keys used to decode the actual file.
// - and then there's also auto delete rule, and some assets type specific stuff.
// - to download relevant resources and/or updating them, the game use the m_asset_package table:
//   - a package is different from a pack:
//     - a package is assosicated with a key, and can contain multiple packs.
//     - and vice versa, a pack can be contained in multiple packages.
//     - pack name are unique, in the sense that if they are the same then they contain the same data no matter what
//   - package_key is a formated string with relevant data for each type of usage.
//     - presumablly, when the game want to access some content, it will generate the relevant package key to look it up.
//   - version is the versioning of the items, can be used to trigger a redownload if some content are updated.
//     - presumablly this is some sort of hash on the file themselves, maybe with salt and stuff.
//     - but it's not actually used to check anything, so any random numbers would do.
//     - pack_num the amount of pack this package key actually contain.
// - after looking at the package key and check for versioning update and suchs, the game would use m_asset_package_mapping to actually download the packs:
//   - package_key is the key of the package:
//     - this a package can contain multiple packs
//     - the number of packs is pack_num as mention above.
//   - pack_name:
//     - this is the name of the relevant pack.
//     - seems to be 6 characters, each can be lowercase Latin or a number
//     - so 36^6 = 2176782336 or enough pack names, although the pack names must be unique.
//     - it's not tested if package with longer / shorter key or other character will work, but they probably will.
//   - file_size:
//     - the file_size of the pack.
//   - metapack_name:
//     - the name of the metapack of this pack, can be null if this pack doesn't have a metapack.
//   - metapack_offset:
//     - where to read from within the metapack, if there is.
//   - category:
//     - this is the asset category, used to calculate the size by type of asset for each type when showing the download dialogue.
// - metapack is the way the game group files usually used togethers, so they would be downloaded togethers.
//   - metapack_name:
//     - this is the name of the metapack, follow the pack name convention.
//     - metapack_name should also be unique along with pack name.
//     - so no pack or metapack can have the same name as another pack or metapack.
//   - file_size and category:
//     - same as pack.
// - generally when downloading, the metapack would be used:
//   - the rule is probably something like this:
//     - whenever a pack is necessary, download that pack.
//     - if that pack is part of a metapack, then download the whole metapack.
//     - however, if the metapack already have something in it, then only download the relevant smaller pack.
//   - this generally work fine and we only need to serve the metapack files + the files that don't have a metapack.
//   - however, in some case, the client still require the file directly:
//     - one way to trigger this is to change the language on global client.
// - so we need to be able to serve a part of the file.
// This package only implement the relevant mapping and lookup, the serving is done either by the external cdn or by our code here.
// We also assume the following:
// - the pack name is different across all locale.
// - if 2 locales share a pack that has the same pack name, then these pack are the same.
import (
	"xorm.io/xorm"
)

var NameToLocale = map[string]string{}
var Metapack = map[string]*MetapackType{}
var Pack = map[string]*PackType{}

func Init(locale string, assetdata *xorm.Engine) {
	session := assetdata.NewSession()
	defer session.Close()
	loadMetapack(locale, session)
	loadPack(locale, session)
}
