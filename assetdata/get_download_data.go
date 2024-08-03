package assetdata

import (
	"fmt"
)

type DownloadData struct {
	Locale       string
	File         string
	IsEntireFile bool // if this is set, the following fields are 0
	Start        int
	Size         int
}

func GetDownloadData(packname string) DownloadData {
	_, exist := Metapack[packname]
	if exist {
		return DownloadData{
			Locale:       NameToLocale[packname],
			File:         packname,
			IsEntireFile: true,
		}
	}
	pack, exist := Pack[packname]
	if !exist {
		panic(fmt.Sprint("package doesn't exist: ", packname))
	}
	if pack.Metapack == nil {
		return DownloadData{
			Locale:       NameToLocale[packname],
			File:         packname,
			IsEntireFile: true,
		}
	} else {
		return DownloadData{
			Locale:       NameToLocale[pack.Metapack.MetapackName],
			File:         pack.Metapack.MetapackName,
			IsEntireFile: false,
			Start:        pack.MetapackOffset,
			Size:         pack.FileSize,
		}
	}
}
