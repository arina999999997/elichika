package image_form

import (
	"elichika/assetdata"
	"elichika/utils"

	hwdecrypt "github.com/arina999999997/gohwdecrypt"

	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"

	"xorm.io/xorm"
)

type Texture struct {
	AssetPath string `xorm:"asset_path"`
	PackName  string `xorm:"pack_name"`
	Head      int    `xorm:"head"`
	Size      int    `xorm:"size"`
	Key1      uint32 `xorm:"key1"`
	Key2      uint32 `xorm:"key2"`
	Width     uint32 `xorm:"width"`
	Height    uint32 `xorm:"height"`
}

var (
	TextureByAssetPath = map[string]Texture{}
	ImageByAssetPath   = map[string]string{}
	MissingAssetImage  string
	UnknownAssetImage  string
)

func init() {
	// texture.db is a database containing the textures of the game
	// if you are interested in playing around with this, contact the relevant people in discord
	engine, err := xorm.NewEngine("sqlite", "texture.db")
	utils.CheckErr(err)
	languages := []string{"en", "ja", "ko", "zh", "th"}
	for _, language := range languages {
		textures := []Texture{}
		err = engine.Table("detailed_texture").Where("language = ? AND error IS NULL", language).OrderBy("time_stamp").Find(&textures)
		utils.CheckErr(err)
		for _, texture := range textures {
			_, exists := TextureByAssetPath[texture.AssetPath]
			if exists {
				continue
			}
			TextureByAssetPath[texture.AssetPath] = texture
		}
	}
}

func getBase64ImageData(data []byte) string {
	return "data:" + http.DetectContentType(data) + ";base64," + base64.StdEncoding.EncodeToString(data)
}

func init() {
	missingAssetFile := "webui/image_form/missing_asset.png"
	stat, err := os.Stat(missingAssetFile)
	utils.CheckErr(err)
	data := make([]byte, stat.Size())
	file, err := os.Open(missingAssetFile)
	utils.CheckErr(err)
	file.Read(data)
	MissingAssetImage = getBase64ImageData(data)
}
func init() {
	unknownAssetFile := "webui/image_form/unknown_asset.png"
	stat, err := os.Stat(unknownAssetFile)
	utils.CheckErr(err)
	data := make([]byte, stat.Size())
	file, err := os.Open(unknownAssetFile)
	utils.CheckErr(err)
	file.Read(data)
	UnknownAssetImage = getBase64ImageData(data)
}

func (t Texture) LoadUnencrypted() (output []byte) {
	output = nil
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	data := assetdata.GetDownloadData(t.PackName)
	actualFile := fmt.Sprintf("static/%s", data.File)
	if _, err := os.Stat(actualFile); errors.Is(err, os.ErrNotExist) {
		return
	}
	actualStart := data.Start + t.Head
	file, err := os.Open(actualFile)
	utils.CheckErr(err)
	defer file.Close()
	output = make([]byte, t.Size)
	_, err = file.ReadAt(output, int64(actualStart))
	utils.CheckErr(err)
	// keying
	hwdecrypt.DecryptBuffer(&hwdecrypt.HwdKeyset{
		Key1: t.Key1,
		Key2: t.Key2,
		Key3: 12345,
	}, output)
	return output
}

func GetImageByAssetPath(assetPath string) string {
	cachedData, cached := ImageByAssetPath[assetPath]
	if cached {
		return cachedData
	}
	texture, known := TextureByAssetPath[assetPath]
	if !known {
		ImageByAssetPath[assetPath] = UnknownAssetImage
		return UnknownAssetImage
	}
	data := texture.LoadUnencrypted()
	if data == nil {
		// doesn't cache the missing as we can potentially find this file and add it
		fmt.Printf("Missing texture: %s (file %s)\n", texture.AssetPath, texture.PackName)
		return MissingAssetImage
	}
	ImageByAssetPath[assetPath] = getBase64ImageData(data)
	return ImageByAssetPath[assetPath]
}

func GetAssetPathsFromSamePackage(assetPath string) []string {
	texture, exists := TextureByAssetPath[assetPath]
	if !exists {
		return nil
	}
	result := []string{}
	for assetPath, otherTexture := range TextureByAssetPath {
		if otherTexture.PackName == texture.PackName {
			result = append(result, assetPath)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return TextureByAssetPath[result[i]].Head < TextureByAssetPath[result[j]].Head
	})
	return result
}
