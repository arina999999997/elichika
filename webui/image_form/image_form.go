package image_form

import (
	"html"
	"strings"
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

func jsEscape(s string) string {
	return html.EscapeString(strings.ReplaceAll(s, `\`, `\\`))
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
	result += `<div><label>` + f.DataLabel + `:</label> <input type="text" id="` + f.DataId + `" name="` + f.DataId + `" autofocus/>`
	result += "\n"
	result += `</form><br>`
	for i, imageName := range f.Keys {
		result += "\n"
		if imageName == "line_break" {
			result += "<br>\n"
			continue
		}
		result += `<a onclick='set_form_data(this, "` + f.DataId + `", "` + jsEscape(imageName) + `")'>`
		result += `<img style="max-height: 31%; max-width: 31%; border: 1cqmin solid #00000000" id="` + jsEscape(imageName) + `" src="` + f.Images[i] + `"/>`
		result += `</a>`
	}
	return result
}
