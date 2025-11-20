package image_form

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/log"
	"elichika/utils"
	"elichika/webui/webui_utils"

	"html"
	// "strings"
	"fmt"

	"xorm.io/xorm"
)

// a form we we can choose the images from
// look something like this:
// - <DataLabel>: <Textbox> to manually input
// - Various images we can choose from:
//   - Clicking on the images will set the text box values
//
// - Images are encoded with base64
type ImageForm struct {
	FormId    string
	DataLabel string
	DataId    string
	Keys      []string
	Images    []string
}

func (f *ImageForm) AddLinebreak() {
	f.Keys = append(f.Keys, "line_break")
	f.Images = append(f.Images, "line_break")
}

func (f *ImageForm) AddImage(imageValue, data string) {
	f.Images = append(f.Images, data)
	f.Keys = append(f.Keys, imageValue)
}

func (f *ImageForm) AddImageByAssetPath(imageValue, assetPath string) {
	data := GetImageByAssetPath(assetPath)
	if imageValue == "" {
		imageValue = assetPath
	}
	f.Images = append(f.Images, data)
	f.Keys = append(f.Keys, imageValue)
}

// TODO(extra): this can't display content amount since that require dynamically rendering stuff
func (f *ImageForm) AddImageByContent(imageValue string, content client.Content) {
	db := gamedata.Instance.MasterdataDb
	assetPath := ""
	switch content.ContentType {
	case enum.ContentTypeCard:
		assetPath, _ = GetCardAssetPaths(content.ContentId)
	case enum.ContentTypeLessonEnhancingItem:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_lesson_enhancing_item").Where("id = ?", content.ContentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeTrainingMaterial:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_training_material").Where("id = ?", content.ContentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeGradeUpper:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_grade_upper").Where("id = ?", content.ContentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeEmblem:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_emblem").Where("id = ?", content.ContentId).Cols("emblem_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeExchangeEventPoint:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_exchange_event_point").Where("id = ?", content.ContentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeLiveSkipTicket:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_live_skip_ticket").Where("id = ?", content.ContentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})

	default:
		log.Panic("unsupported content type")
	}
	data := GetImageByAssetPath(assetPath)
	if imageValue == "" {
		imageValue = assetPath
	}
	f.Images = append(f.Images, data)
	f.Keys = append(f.Keys, imageValue)
}

func (f *ImageForm) AddImageByAssetPathFilterBySize(imageValue, assetPath string, width, height uint32) {
	data := GetImageByAssetPath(assetPath)
	texture, exist := TextureByAssetPath[assetPath]
	if !exist {
		panic("Asset path doesn't have a texture, cannot be used for filter by size: " + assetPath)
	}
	if (texture.Width != width) || (texture.Height != height) {
		return
	}
	if imageValue == "" {
		imageValue = assetPath
	}
	f.Images = append(f.Images, data)
	f.Keys = append(f.Keys, imageValue)
}

func (f *ImageForm) GetHTML() string {
	result := `<form id="` + f.FormId + `" method="POST" enctype="multipart/form-data">`
	result += "\n"
	result += `<div><label>` + f.DataLabel + `:</label> <input type="text" id="` + f.DataId + `" name="` + f.DataId + `" autofocus/> </div>`
	result += "\n"
	result += `</form><br>`
	result += "<div class=\"form-row\">\n"
	for i, imageName := range f.Keys {
		result += "\n"
		if imageName == "line_break" {
			result += "</div>\n"
			result += "<div class=\"form-row\">\n"
			continue
		}
		script := fmt.Sprintf(`set_form_data(this, %s, %s)`, webui_utils.JsString(f.DataId), webui_utils.JsString(imageName))
		result += fmt.Sprintf(`<a onclick="%s">`, html.EscapeString(script))
		result += `<div class="form-option">`
		result += `<img class="form-image-option" src="` + f.Images[i] + `"/>`
		result += `</div></a>`
	}
	result += "</div>\n"
	return result
}
