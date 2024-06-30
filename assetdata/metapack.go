package assetdata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MetapackType struct {
	MetapackName string `xorm:"pk 'metapack_name'"`
	FileSize     int    `xorm:"'file_size'"`
	Category     int    `xorm:"'category'"`
}

func loadMetapack(locale string, session *xorm.Session) {
	fmt.Println("Loading Metapack")
	metapacks := map[string]*MetapackType{}
	err := session.Table("metapack").Find(&metapacks)
	utils.CheckErr(err)
	for name, metapack := range metapacks {
		previous, exist := Metapack[name]
		if !exist {
			Metapack[name] = metapack
			NameToLocale[name] = locale
			continue
		}
		if (previous.FileSize != metapack.FileSize) || (previous.Category != metapack.Category) {
			panic(fmt.Sprint("Metapack name reused: ", *previous, *metapack, "\nLocale: ", locale, ", previous locale: ", NameToLocale[name]))
		}
	}
}
