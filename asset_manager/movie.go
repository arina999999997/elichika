package asset_manager

import (
	"elichika/log"

	"fmt"
)

type Movie struct {
	Pavement string `xorm:"'pavement'"`
	PackName string `xorm:"'pack_name'"`

	Locale string `xorm:"-"`
}

/*
IMPLEMENTING RAW ASSET INTERFACE
*/
func (m Movie) GetLocale() string {
	if m.Locale != "" {
		return m.Locale
	}
	return "un" // un locale for universal
}

func (m Movie) GetAssetPath() string {
	return m.Pavement
}

func (m Movie) GetUpdateSQL(assetDatabase *AssetDatabase) string {
	existingMovie, exist := AssetDatabaseGetMovie(assetDatabase, m.Pavement)
	if !exist {
		return fmt.Sprintf("INSERT INTO m_movie VALUES (\"%s\", \"%s\");\n",
			sqlEscape(m.Pavement), m.PackName)
	}
	if m.Equal(existingMovie) {
		return "" // no update necessary
	}
	return fmt.Sprintf("UPDATE m_movie SET pack_name = \"%s\" WHERE pavement = \"%s\";\n",
		m.PackName, sqlEscape(m.Pavement))
}

func (m Movie) GetPack() *Pack {
	return GetPack(m.PackName)
}

func (m Movie) GetAwbPack() *Pack {
	return nil
}

func (m Movie) IsAvailable() bool {
	return m.GetPack().IsAvailable()
}

func (m Movie) Info() string {
	return fmt.Sprintf("movie_%s(%s,%s)", m.GetLocale(), m.Pavement, m.PackName)
}

func (m Movie) InsertIntoAssetDatabase(assetDatabase *AssetDatabase) {
	_, exist := assetDatabase.MovieMap[m.Pavement]
	if exist {
		panic("Asset database already have pavement: " + m.Pavement)
	}
	assetDatabase.Movies = append(assetDatabase.Movies, m)
	assetDatabase.MovieMap[m.Pavement] = m
}

func (m Movie) Equal(other Movie) bool {
	if m.Pavement != other.Pavement {
		return false
	}
	if m.PackName != other.PackName {
		log.Printf("Unexpected movie asset mismatch: %s vs %s", m.Info(), other.Info())
		return false
	}
	return true
}
