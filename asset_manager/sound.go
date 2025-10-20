package asset_manager

import (
	"elichika/log"

	"fmt"
)

type Sound struct {
	SheetName   string `xorm:"'sheet_name'"`
	AcbPackName string `xorm:"'acb_pack_name'"`
	AwbPackName string `xorm:"'awb_pack_name'"`

	Locale string `xorm:"-"`
}

/*
IMPLEMENTING RAW ASSET INTERFACE
*/
func (s Sound) GetLocale() string {
	if s.Locale != "" {
		return s.Locale
	}
	return "un" // un locale for universal
}

func (s Sound) GetAssetPath() string {
	return s.SheetName
}

func (s Sound) GetUpdateSQL(assetDatabase *AssetDatabase) string {
	existingSound, exist := AssetDatabaseGetSound(assetDatabase, s.SheetName)
	if !exist {
		return fmt.Sprintf("INSERT INTO m_asset_sound VALUES (\"%s\", \"%s\", \"%s\");\n",
			sqlEscape(s.SheetName), s.AcbPackName, s.AwbPackName)
	}
	if s.Equal(existingSound) {
		return "" // no update necessary
	}
	return fmt.Sprintf("UPDATE m_asset_sound SET acb_pack_name = \"%s\", awb_pack_name = \"%s\" WHERE sheet_name = \"%s\";\n",
		s.AcbPackName, s.AwbPackName, sqlEscape(s.SheetName))
}

func (s Sound) GetPack() *Pack {
	return GetPack(s.AcbPackName)
}

func (s Sound) GetAwbPack() *Pack {
	return GetPack(s.AwbPackName)
}

func (s Sound) IsAvailable() bool {
	return s.GetPack().IsAvailable() && s.GetAwbPack().IsAvailable()
}

func (s Sound) Info() string {
	return fmt.Sprintf("sound_%s(%s,%s,%s)", s.GetLocale(), s.SheetName, s.AcbPackName, s.AwbPackName)
}

func (s Sound) InsertIntoAssetDatabase(assetDatabase *AssetDatabase) {
	_, exist := assetDatabase.SoundMap[s.SheetName]
	if exist {
		panic("Asset database already have sheet name: " + s.SheetName)
	}
	assetDatabase.Sounds = append(assetDatabase.Sounds, s)
	assetDatabase.SoundMap[s.SheetName] = s
}

func (s Sound) Equal(other Sound) bool {
	if s.SheetName != other.SheetName {
		return false
	}
	if (s.AcbPackName != other.AcbPackName) || (s.AwbPackName != other.AwbPackName) {
		log.Printf("Unexpected sound asset mismatch: %s vs %s", s.Info(), other.Info())
		return false
	}
	return true
}
