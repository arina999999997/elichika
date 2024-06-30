package assetdata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type PackType struct {
	PackageKey     string        `xorm:"'package_key'"`
	PackName       string        `xorm:"'pack_name'"`
	FileSize       int           `xorm:"'file_size'"`
	MetapackName   *string       `xorm:"'metapack_name'"`
	Metapack       *MetapackType `xorm:"-"`
	MetapackOffset int           `xorm:"'metapack_offset'"`
	Category       int           `xorm:"'category'"`
}

func loadPack(locale string, session *xorm.Session) {
	fmt.Println("Loading Pack")
	packs := []*PackType{}
	err := session.Table("m_asset_package_mapping").Find(&packs)
	utils.CheckErr(err)
	for _, pack := range packs {
		_, exist := Metapack[pack.PackName]
		if exist {
			panic(fmt.Sprint("same name used for both a pack and a metapack: ", pack.PackName,
				"\nLocale: ", locale, ", previous locale: ", NameToLocale[pack.PackName]))
		}
		if pack.MetapackName != nil {
			pack.Metapack = Metapack[*pack.MetapackName]
		}
		previous, exist := Pack[pack.PackName]
		if !exist {
			Pack[pack.PackName] = pack
			NameToLocale[pack.PackName] = locale
			continue
		}
		if (previous.FileSize != pack.FileSize) || (previous.Category != pack.Category) {
			// TODO(assert): This doesn't necessarily imply all packs are the same, some test is necessary
			panic(fmt.Sprint("pack name reused: ", *previous, *pack))
		}
	}
}
