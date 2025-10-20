package asset_manager

// importer imports assets into the asset manager
// for now this is done by performing load operation and calling directly into RegisterRawAsset and RegisterPack
// Some possible importer:
// - importer_db.go:
//   - import from an existing asset database file (i.e. asset_a_en.db)
//   - this format is also used for the cached database too
// - importer_git_diff.go (TODO)
//   - import the asset information from database diff (github.com/kotori2/-llasww_master_db_diff.git and the jp version)
//   - this might not be necessary, we can just parse all the data from various archive database
//   - but it might be faster

import (
	"elichika/log"
	"elichika/utils"

	"fmt"
	"sort"
	"strings"
	"time"

	"xorm.io/xorm"
)

func ImportFromAssetsRepository() {
	ImportAssetFromDatabase("assets/db/gl/asset_a_en.db", "en")
	ImportAssetFromDatabase("assets/db/gl/asset_i_en.db", "en")
	ImportAssetFromDatabase("assets/db/gl/asset_a_ko.db", "ko")
	ImportAssetFromDatabase("assets/db/gl/asset_i_ko.db", "ko")
	ImportAssetFromDatabase("assets/db/gl/asset_a_zh.db", "zh")
	ImportAssetFromDatabase("assets/db/gl/asset_i_zh.db", "zh")
	ImportAssetFromDatabase("assets/db/jp/asset_a_ja.db", "ja")
	ImportAssetFromDatabase("assets/db/jp/asset_i_jp.db", "ja")
}

func ImportCachedAssets() {
	if utils.PathExists("asset_manager_cached_asset_ja.db") {
		ImportAssetFromDatabase("asset_manager_cached_asset_ja.db", "ja")
	}
	if utils.PathExists("asset_manager_cached_asset_en.db") {
		ImportAssetFromDatabase("asset_manager_cached_asset_en.db", "en")
	}
	if utils.PathExists("asset_manager_cached_asset_ko.db") {
		ImportAssetFromDatabase("asset_manager_cached_asset_ko.db", "ko")
	}
	if utils.PathExists("asset_manager_cached_asset_zh.db") {
		ImportAssetFromDatabase("asset_manager_cached_asset_zh.db", "zh")
	}
}

// import from a directory that have the following structures:
// - version.db with version table that have 3 rows:
//   - region, hash, time
//   - region is en or jp for ww or jp
//   - hash is the hash version
//   - time is a time string with format yyyy-MMM-dd hh:mm
//
// - then for each version we have a directory <version jp/en>-<hash>
//   - inside it there's asset_i_ja.db or asset_i_en.db or other version
//
// - parse in reverse order of time
// note:
// - if you want this data, contact triangle or search the relevant channels for downloading tools
// - it's up to you to format the database and create the version database
// due to how long this take, we will split this into chunk and load from the cached and write to cache in different run
// so this take a first and last limit
func ImportFromArchiveDatabases(dir, minTime, maxTime string) {
	if (dir != "") && (dir[len(dir)-1] != '/') {
		dir += "/"
	}
	engine, err := xorm.NewEngine("sqlite", dir+"version.db")
	utils.CheckErr(err)
	defer engine.Close()
	type Version struct {
		Region   string `xorm:"'region'"`
		Hash     string `xorm:"'hash'"`
		Time     string `xorm:"'time'"`
		UnixTime int64  `xorm:"-"`
	}
	versions := []Version{}
	err = engine.Table("version").Find(&versions)
	utils.CheckErr(err)
	const layout = "2006-Jan-02 15:04"
	for i := range versions {
		timePoint, err := time.Parse(layout, strings.Trim(versions[i].Time, "\n"))
		utils.CheckErr(err)
		versions[i].UnixTime = timePoint.Unix()
	}
	unixMin, err := time.Parse(layout, minTime)
	utils.CheckErr(err)
	unixMax, err := time.Parse(layout, maxTime)
	utils.CheckErr(err)
	if unixMin.Unix() > unixMax.Unix() {
		panic("minTime must be before or equal maxTime")
	}
	log.Printf("Processing archive database between %s and %s\n", minTime, maxTime)

	sort.Slice(versions, func(i, j int) bool {
		if versions[i].UnixTime != versions[j].UnixTime {
			return versions[i].UnixTime > versions[j].UnixTime
		}
		if versions[i].Region == versions[j].Region {
			panic("Same region and time")
		}
		return versions[i].Region == "jp"
	})
	for _, version := range versions {
		if (version.UnixTime < unixMin.Unix()) || (version.UnixTime > unixMax.Unix()) {
			continue
		}
		log.Printf("Loading from archive database, Region: %s, Time: %s, Hash: %s\n", version.Region, strings.Trim(version.Time, "\n"), version.Hash)
		versionDir := fmt.Sprintf("%s%s-%s/", dir, version.Region, version.Hash)
		if version.Region == "en" {
			if utils.PathExists(versionDir + "asset_i_en.db") {
				ImportAssetFromDatabase(versionDir+"asset_i_en.db", "en")
			} else {
				log.Printf("WARNING: No en database\n")
			}
			if utils.PathExists(versionDir + "asset_i_ko.db") {
				ImportAssetFromDatabase(versionDir+"asset_i_ko.db", "ko")
			} else {
				log.Printf("WARNING: No ko database\n")
			}
			if utils.PathExists(versionDir + "asset_i_zh.db") {
				ImportAssetFromDatabase(versionDir+"asset_i_zh.db", "zh")
			} else {
				log.Printf("WARNING: No zh database\n")
			}
		} else if version.Region == "jp" {
			if utils.PathExists(versionDir + "asset_i_ja.db") {
				ImportAssetFromDatabase(versionDir+"asset_i_ja.db", "ja")
			} else if utils.PathExists(versionDir + "asset_i_ja_0.db") {
				ImportAssetFromDatabase(versionDir+"asset_i_ja_0.db", "ja")
			} else {
				log.Printf("WARNING: No ja database\n")
			}
		} else {
			panic("Unknown region: " + version.Region)
		}
	}
}
