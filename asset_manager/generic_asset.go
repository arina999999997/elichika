package asset_manager

// with the exception of sound and movie, every assets have this structure
import (
	"elichika/log"

	"fmt"
)

type AssetTypeSpecifier interface {
	TableName() string
}

type GenericAsset[Specifier AssetTypeSpecifier] struct {
	AssetPath string `xorm:"pk 'asset_path'"`
	PackName  string `xorm:"'pack_name'"`
	Head      int32  `xorm:"'head'"`
	Size      int32  `xorm:"'size'"`
	Key1      int32  `xorm:"'key1'"`
	Key2      int32  `xorm:"'key2'"`
	// extra information about locale of the file, might not be filled at all
	Locale string `xorm:"-"`
}

/*
IMPLEMENTING RAW ASSET INTERFACE
*/
func (a GenericAsset[Specifier]) GetLocale() string {
	return a.Locale
}

func (a GenericAsset[Specifier]) GetAssetPath() string {
	return a.AssetPath
}

func (a GenericAsset[Specifier]) GetUpdateSQL(assetDatabase *AssetDatabase) string {
	var specifier Specifier
	existingAsset, exist := AssetDatabaseGetGenericAsset[Specifier](assetDatabase, a.AssetPath)
	if !exist {
		return fmt.Sprintf("INSERT INTO %s VALUES (\"%s\", \"%s\", %d, %d, %d, %d);\n",
			specifier.TableName(), sqlEscape(a.AssetPath), a.PackName, a.Head, a.Size, a.Key1, a.Key2)
	}
	if existingAsset.Equal(a) {
		return "" // no update necessary
	}
	return fmt.Sprintf("UPDATE %s SET pack_name = \"%s\", head = %d, size = %d, key1 = %d, key2 = %d WHERE asset_path = \"%s\";\n",
		specifier.TableName(), a.PackName, a.Head, a.Size, a.Key1, a.Key2, sqlEscape(a.AssetPath))
}

func (a GenericAsset[Specifier]) GetPack() *Pack {
	return GetPack(a.PackName)
}

func (a GenericAsset[Specifier]) GetAwbPack() *Pack {
	return nil
}

func (a GenericAsset[Specifier]) IsAvailable() bool {
	return a.GetPack().IsAvailable()
}

func (a GenericAsset[Specifier]) Info() string {
	// key is ommited as they're usually not important, and they can appear in the wrong order too
	var specifier Specifier
	return fmt.Sprintf("%s_%s(%s,%s,%d,%d)", specifier.TableName(), a.GetLocale(), a.AssetPath, a.PackName, a.Head, a.Size)
}

func (a GenericAsset[Specifier]) InsertIntoAssetDatabase(assetDatabase *AssetDatabase) {
	var specifier Specifier
	tableName := specifier.TableName()
	assetDatabase.GenericAssetsByTableName[tableName] = append(assetDatabase.GenericAssetsByTableName[tableName].([]GenericAsset[Specifier]), a)
	_, exists := assetDatabase.GenericAssetMapByTableName[tableName].(map[string]GenericAsset[Specifier])[a.AssetPath]
	if exists {
		panic("asset database already have asset path: " + a.AssetPath)
	}
	assetDatabase.GenericAssetMapByTableName[tableName].(map[string]GenericAsset[Specifier])[a.AssetPath] = a
}

/*
  SPEICFIC TYPED FUNCTIONS
*/

func (a GenericAsset[Specifier]) Equal(other GenericAsset[Specifier]) bool {
	if a.AssetPath != other.AssetPath {
		return false
	}
	if a.PackName != other.PackName {
		return false
	}
	if (a.Head != other.Head) || (a.Size != other.Size) || (a.Key1 != other.Key1) || (a.Key2 != other.Key2) { // this shouldn't happen so we log them out just in case
		log.Printf("Unexpected generic asset mismatch: %s vs %s", a.Info(), other.Info())
		return false
	}
	return true
}

/*
  SPECIFIERS
  Note that sound and movie are handled differently from the generic
*/

type AdvScriptSpecifier struct {
}

func (_ AdvScriptSpecifier) TableName() string {
	return "adv_script"
}

type BackgroundSpecifier struct {
}

func (_ BackgroundSpecifier) TableName() string {
	return "background"
}

type GachaPerformanceSpecifier struct {
}

func (_ GachaPerformanceSpecifier) TableName() string {
	return "gacha_performance"
}

type Live2dSdModelSpecifier struct {
}

func (_ Live2dSdModelSpecifier) TableName() string {
	return "live2d_sd_model"
}

type LivePropSkeletonSpecifier struct {
}

func (_ LivePropSkeletonSpecifier) TableName() string {
	return "live_prop_skeleton"
}

type LiveTimelineSpecifier struct {
}

func (_ LiveTimelineSpecifier) TableName() string {
	return "live_timeline"
}

type MemberFacialSpecifier struct {
}

func (_ MemberFacialSpecifier) TableName() string {
	return "member_facial"
}

type MemberFacialAnimationSpecifier struct {
}

func (_ MemberFacialAnimationSpecifier) TableName() string {
	return "member_facial_animation"
}

type MemberModelSpecifier struct {
}

func (_ MemberModelSpecifier) TableName() string {
	return "member_model"
}

type MemberSdModelSpecifier struct {
}

func (_ MemberSdModelSpecifier) TableName() string {
	return "member_sd_model"
}

type NaviMotionSpecifier struct {
}

func (_ NaviMotionSpecifier) TableName() string {
	return "navi_motion"
}

type NaviTimelineSpecifier struct {
}

func (_ NaviTimelineSpecifier) TableName() string {
	return "navi_timeline"
}

type ShaderSpecifier struct {
}

func (_ ShaderSpecifier) TableName() string {
	return "shader"
}

type SkillEffectSpecifier struct {
}

func (_ SkillEffectSpecifier) TableName() string {
	return "skill_effect"
}

type SkillTimelineSpecifier struct {
}

func (_ SkillTimelineSpecifier) TableName() string {
	return "skill_timeline"
}

type SkillWipeSpecifier struct {
}

func (_ SkillWipeSpecifier) TableName() string {
	return "skill_wipe"
}

type StageSpecifier struct {
}

func (_ StageSpecifier) TableName() string {
	return "stage"
}

type StageEffectSpecifier struct {
}

func (_ StageEffectSpecifier) TableName() string {
	return "stage_effect"
}

type TextureSpecifier struct {
}

func (_ TextureSpecifier) TableName() string {
	return "texture"
}
