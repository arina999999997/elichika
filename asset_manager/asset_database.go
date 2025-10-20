package asset_manager

import (
	"elichika/log"
	"elichika/utils"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

// represent an existing asset database
// to improve speed we load these into memory and perform operation on them instead of calling into the db everytime
type AssetDatabase struct {
	// Maybe some version info
	Engine *xorm.Engine

	// map from table name to slice
	GenericAssetsByTableName map[string]any
	// map from table name to object inside above slide
	GenericAssetMapByTableName map[string]any

	Sounds   []Sound
	SoundMap map[string]Sound
	Movies   []Movie
	MovieMap map[string]Movie

	Packs             []Pack          // list of pack
	PackMap           map[string]bool // whether a pack exist
	PackIsMainPackage map[string]bool // whether a pack is already included in main package

	Importing bool // opened for importing the data
	Exporting bool // opened for exporting the data (for cache)
	Updating  bool // opened for generating sql update (for releasing)
}

func NewAssetDatabaseForImporting(path string) *AssetDatabase {
	assetDatabase := &AssetDatabase{
		Importing: true,
	}
	var err error
	assetDatabase.Engine, err = xorm.NewEngine("sqlite", path)
	utils.CheckErr(err)
	assetDatabase.Load()
	return assetDatabase
}

// create an asset database for writing purpose
// the database at path should be empty, Load is still called to create the relevant tables
func NewAssetDatabaseForExporting(path string) *AssetDatabase {
	assetDatabase := &AssetDatabase{
		Exporting: true,
	}
	var err error
	assetDatabase.Engine, err = xorm.NewEngine("sqlite", path)
	utils.CheckErr(err)
	assetDatabase.Load()
	return assetDatabase
}

func NewAssetDatabaseForUpdating(path string) *AssetDatabase {
	assetDatabase := &AssetDatabase{
		Updating: true,
	}
	var err error
	assetDatabase.Engine, err = xorm.NewEngine("sqlite", path)
	utils.CheckErr(err)
	assetDatabase.Load()
	return assetDatabase
}

// note that this only close the db file
// the memory stay around until no more reference to it exists
func FreeAssetDatabase(assetDatabase *AssetDatabase) {
	err := assetDatabase.Engine.Close()
	utils.CheckErr(err)
	// explicitly free other things too, they should be copied into relevant place by now
	assetDatabase.GenericAssetsByTableName = nil
	assetDatabase.GenericAssetMapByTableName = nil
	assetDatabase.Sounds = nil
	assetDatabase.SoundMap = nil
	assetDatabase.Movies = nil
	assetDatabase.MovieMap = nil
	assetDatabase.Packs = nil
	assetDatabase.PackMap = nil
	assetDatabase.PackIsMainPackage = nil
}

func (assetDatabase *AssetDatabase) LoadPacks() {
	if assetDatabase.Updating {
		// if we are updating, we need to load the full data with package key
		type PackWithPackageKey struct {
			PackageKey     string  `xorm:"'package_key'"`
			PackName       string  `xorm:"'pack_name'"`
			FileSize       int32   `xorm:"'file_size'"`
			MetapackName   *string `xorm:"'metapack_name'"`
			MetapackOffset int32   `xorm:"'metapack_offset'"`
		}
		packs := []PackWithPackageKey{}
		err := assetDatabase.Engine.Table("m_asset_package_mapping").Find(&packs)
		utils.CheckErr(err)
		assetDatabase.PackMap = map[string]bool{}
		assetDatabase.PackIsMainPackage = map[string]bool{}
		for _, pack := range packs {
			if pack.PackageKey == "main" {
				assetDatabase.PackIsMainPackage[pack.PackName] = true
			}
			if assetDatabase.PackMap[pack.PackName] {
				continue
			}
			assetDatabase.PackMap[pack.PackName] = true
			assetDatabase.Packs = append(assetDatabase.Packs, Pack{
				PackName:       pack.PackName,
				FileSize:       pack.FileSize,
				MetapackName:   pack.MetapackName,
				MetapackOffset: pack.MetapackOffset,
			})
		}
	} else {
		// otherwise the package key is not necessary
		err := assetDatabase.Engine.Table("m_asset_package_mapping").Find(&assetDatabase.Packs)
		utils.CheckErr(err)
	}
}

func (assetDatabase *AssetDatabase) Load() {
	// first load the pack for the loaders to use
	assetDatabase.LoadPacks()
	// create the relevant maps
	assetDatabase.GenericAssetsByTableName = map[string]any{}
	assetDatabase.GenericAssetMapByTableName = map[string]any{}
	// call load with the specifier
	AssetDatabaseLoadGenericAsset[AdvScriptSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[BackgroundSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[GachaPerformanceSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[Live2dSdModelSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[LivePropSkeletonSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[LiveTimelineSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[MemberFacialSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[MemberFacialAnimationSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[MemberModelSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[MemberSdModelSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[NaviMotionSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[NaviTimelineSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[ShaderSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[SkillEffectSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[SkillTimelineSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[SkillWipeSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[StageSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[StageEffectSpecifier](assetDatabase)
	AssetDatabaseLoadGenericAsset[TextureSpecifier](assetDatabase)
	// load sound and movies
	AssetDatabaseLoadSound(assetDatabase)
	AssetDatabaseLoadMovie(assetDatabase)
}

func AssetDatabaseLoadGenericAsset[Specifier AssetTypeSpecifier](assetDatabase *AssetDatabase) {
	var specifier Specifier
	tableName := specifier.TableName()
	slice := []GenericAsset[Specifier]{}
	exist, err := assetDatabase.Engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)
	if !exist {
		return
	}
	err = assetDatabase.Engine.Table(tableName).Find(&slice)
	utils.CheckErr(err)
	assetMap := map[string]GenericAsset[Specifier]{}
	for i := range slice {
		assetMap[slice[i].AssetPath] = slice[i]
	}
	assetDatabase.GenericAssetsByTableName[tableName] = slice
	assetDatabase.GenericAssetMapByTableName[tableName] = assetMap
}

func AssetDatabaseGetGenericAsset[Specifier AssetTypeSpecifier](assetDatabase *AssetDatabase, assetPath string) (GenericAsset[Specifier], bool) {
	var specifier Specifier
	asset, exist := assetDatabase.GenericAssetMapByTableName[specifier.TableName()].(map[string]GenericAsset[Specifier])[assetPath]
	return asset, exist
}

func AssetDatabaseLoadSound(assetDatabase *AssetDatabase) {
	err := assetDatabase.Engine.Table("m_asset_sound").Find(&assetDatabase.Sounds)
	utils.CheckErr(err)
	assetDatabase.SoundMap = map[string]Sound{}
	for i := range assetDatabase.Sounds {
		assetDatabase.SoundMap[assetDatabase.Sounds[i].SheetName] = assetDatabase.Sounds[i]
	}
}

func AssetDatabaseGetSound(assetDatabase *AssetDatabase, sheetName string) (Sound, bool) {
	sound, exist := assetDatabase.SoundMap[sheetName]
	return sound, exist
}

func AssetDatabaseLoadMovie(assetDatabase *AssetDatabase) {
	exist, err := assetDatabase.Engine.Table("m_movie").IsTableExist("m_movie")
	utils.CheckErr(err)
	if !exist {
		return
	}
	err = assetDatabase.Engine.Table("m_movie").Find(&assetDatabase.Movies)
	utils.CheckErr(err)
	assetDatabase.MovieMap = map[string]Movie{}
	for i := range assetDatabase.Movies {
		assetDatabase.MovieMap[assetDatabase.Movies[i].Pavement] = assetDatabase.Movies[i]
	}
}

func AssetDatabaseGetMovie(assetDatabase *AssetDatabase, pavement string) (Movie, bool) {
	movie, exist := assetDatabase.MovieMap[pavement]
	return movie, exist
}

func AssetDatabaseExportToDbGenericAsset[Specifier AssetTypeSpecifier](assetDatabase *AssetDatabase) {
	var specifier Specifier
	tableName := specifier.TableName()
	slice := assetDatabase.GenericAssetsByTableName[tableName].([]GenericAsset[Specifier])
	n := len(slice)
	log.Printf("\tInserting into %s\n", tableName)
	for start := 0; start < n; start += 1024 {
		end := start + 1024
		if end > n {
			end = n
		}
		inserted, err := assetDatabase.Engine.Table(tableName).Insert(slice[start:end])
		utils.CheckErr(err)
		if int(inserted) != end-start {
			panic("Failed to insert properly")
		}
	}
}

func (assetDatabase *AssetDatabase) ExportToDb() {
	if !assetDatabase.Exporting {
		panic("Cannot write to database not marked as writing")
	}
	log.Println("Writing database")
	AssetDatabaseExportToDbGenericAsset[AdvScriptSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[BackgroundSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[GachaPerformanceSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[Live2dSdModelSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[LivePropSkeletonSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[LiveTimelineSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[MemberFacialSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[MemberFacialAnimationSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[MemberModelSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[MemberSdModelSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[NaviMotionSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[NaviTimelineSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[ShaderSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[SkillEffectSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[SkillTimelineSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[SkillWipeSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[StageSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[StageEffectSpecifier](assetDatabase)
	AssetDatabaseExportToDbGenericAsset[TextureSpecifier](assetDatabase)
	{
		n := len(assetDatabase.Sounds)
		log.Printf("\tInserting into m_asset_sound\n")
		for start := 0; start < n; start += 1024 {
			end := start + 1024
			if end > n {
				end = n
			}
			inserted, err := assetDatabase.Engine.Table("m_asset_sound").Insert(assetDatabase.Sounds[start:end])
			utils.CheckErr(err)
			if int(inserted) != end-start {
				panic("Failed to insert properly")
			}
		}
	}
	{
		n := len(assetDatabase.Movies)
		log.Printf("\tInserting into m_movie\n")
		for start := 0; start < n; start += 1024 {
			end := start + 1024
			if end > n {
				end = n
			}
			inserted, err := assetDatabase.Engine.Table("m_movie").Insert(assetDatabase.Movies[start:end])
			utils.CheckErr(err)
			if int(inserted) != end-start {
				panic("Failed to insert properly")
			}
		}
	}
	{
		n := len(Packs)
		log.Printf("\tInserting into m_asset_package_mapping\n")
		for start := 0; start < n; start += 1024 {
			end := start + 1024
			if end > n {
				end = n
			}
			inserted, err := assetDatabase.Engine.Table("m_asset_package_mapping").Insert(Packs[start:end])
			utils.CheckErr(err)
			if int(inserted) != end-start {
				panic("Failed to pack properly")
			}
		}
	}
}
