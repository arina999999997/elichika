package asset_manager

import (
	"elichika/log"
	"elichika/utils"

	"fmt"
)

// represent a Pack as the game see it
type Pack struct {
	// the data as the game see it from m_asset_package_mapping
	PackName       string  `xorm:"'pack_name'"`
	FileSize       int32   `xorm:"'file_size'"`
	MetapackName   *string `xorm:"'metapack_name'"`
	MetapackOffset int32   `xorm:"'metapack_offset'"`
	// Category       int32   `xorm:"'category'"` // early data doesn't have this part
	// extra info
	HasLocalFile *bool `xorm:"-"`
}

func (pack *Pack) IsAvailable() bool {
	if pack.HasLocalFile != nil {
		return *pack.HasLocalFile
	}
	pack.HasLocalFile = new(bool)
	if pack.MetapackName != nil {
		*pack.HasLocalFile = utils.PathExists(fmt.Sprintf("static/%s", *pack.MetapackName))
		if !*pack.HasLocalFile {
			*pack.HasLocalFile = utils.PathExists(fmt.Sprintf("static/%s", pack.PackName))
			if *pack.HasLocalFile {
				log.Printf("WARNING: Metapack doesn't exists but pack exist: pack: %s metapack: %s\n", pack.PackName, *pack.MetapackName)
				log.Printf("WARNING: Removed the metapack attribute from the pack")
				pack.MetapackName = nil
				pack.MetapackOffset = 0
			}
		}
	} else {
		*pack.HasLocalFile = utils.PathExists(fmt.Sprintf("static/%s", pack.PackName))
	}
	return *pack.HasLocalFile
}

func (pack *Pack) GetUpdateSQL(assetDatabase *AssetDatabase) string {
	if !assetDatabase.Updating {
		panic("asset database is not set to updating")
	}
	if assetDatabase.PackIsMainPackage[pack.PackName] { // already exist and included in the main package (downloaded by default), no need to insert it
		return ""
	}
	if assetDatabase.PackMap[pack.PackName] {
		// exist but not inside main, so we insert the main package key for this pack to be available
		// TODO(extra): maybe it's necessary to set the autodelete flag too, but probably not
		// TODO(extra): set the category properly instead of 1
		return fmt.Sprintf("INSERT INTO m_asset_package_mapping VALUES(\"main\", \"%s\", %d, %s, %d, %d);\n",
			pack.PackName, pack.FileSize, sqlNullString(pack.MetapackName), pack.MetapackOffset, 1)
	}
	// otherwise we need to both insert and update the package map
	sql := fmt.Sprintf("INSERT INTO m_asset_pack VALUES(\"%s\", 0);\n", pack.PackName)
	sql += fmt.Sprintf("INSERT INTO m_asset_package_mapping VALUES(\"main\", \"%s\", %d, %s, %d, %d);\n",
		pack.PackName, pack.FileSize, sqlNullString(pack.MetapackName), pack.MetapackOffset, 1)
	// TODO(extra): Handle metapack whenever we actually run into it
	if pack.MetapackName != nil {
		panic("TODO: Handle metapack for pack updating")
	}
	return sql
}

func (pack *Pack) GetCdnCheck(assetDatabase *AssetDatabase, packs *[]string) {
	if !assetDatabase.Updating {
		panic("asset database is not set to updating")
	}
	if assetDatabase.PackIsMainPackage[pack.PackName] { // already exist and included in the main package (downloaded by default), no need to insert it
		return
	}
	if assetDatabase.PackMap[pack.PackName] {
		// exist but not inside main, so we insert the main package key for this pack to be available
		return
	}
	if pack.MetapackName != nil {
		*packs = append(*packs, *pack.MetapackName)
	} else {
		*packs = append(*packs, pack.PackName)
	}
}
