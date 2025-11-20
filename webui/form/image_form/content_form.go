package image_form

import (
	"elichika/webui/webui_utils"

	"fmt"
	"html"
)

// content form is a form where the options are related to the content of the game
type ContentForm struct {
	FormId            string
	DataLabel         string
	DataId            string
	Rows              [][]string
	RowCount          int
	TradeToggleButton bool
}

func (f *ContentForm) NewRow() {
	f.Rows = append(f.Rows, []string{})
	f.RowCount += 1
}

const formOptionHTMLFormatScriptString = `set_form_data(this, %s, %s)`
const formOptionHTMLFormatString = `<a onclick="%s"> <div class="form-option">%s</div></a>`

func (f *ContentForm) AddImage(value, imageHTML string) {
	if f.RowCount == 0 {
		f.NewRow()
	}
	script := fmt.Sprintf(formOptionHTMLFormatScriptString, webui_utils.JsString(f.DataId), webui_utils.JsString(value))
	f.Rows[f.RowCount-1] = append(f.Rows[f.RowCount-1], fmt.Sprintf(formOptionHTMLFormatString, html.EscapeString(script), imageHTML))
}

// TODO(webui visual): helper functions
// func (f *ContentForm) AddImageByAssetPath(imageValue, assetPath string) {
// }

// func (f *ContentForm) AddImageByContent(imageValue string, content client.Content) {
// }

// func (f *ContentForm) AddImageByAssetPathFilterBySize(imageValue, assetPath string, width, height uint32) {
// }

func (f *ContentForm) GetHTML() string {
	// result := `<form id="` + f.FormId + `" method="POST" enctype="multipart/form-data">`
	result := fmt.Sprintf(`<form id=%s method="POST" enctype="multipart/form-data">`, webui_utils.JsString(f.FormId))
	result += "\n"
	result += fmt.Sprintf(`<div><label>%s</label> <input type="text" id=%s name=%s autofocus /> </div>`, f.DataLabel, webui_utils.JsString(f.DataId), webui_utils.JsString(f.DataId))
	result += "\n"
	result += `</form><br>`
	// add a button that toggle between avaiable and unavailable shop item
	if f.TradeToggleButton {
		result +=
			`<script>
function toggle_trade() {
    items = document.getElementsByClassName("trade-item-unavailable")
	if (items.length > 0) {
    	while (items.length > 0) {
			items[0].className = "trade-item"
		}
		return
	}
	items = document.getElementsByClassName("trade-item")
	for (i = 0; i < items.length; i++) {
		items[i].className = "trade-item trade-item-unavailable"
	}
}
</script>
<div>
<button onclick="toggle_trade()"> Toggle Trade </button>
</div>
`
	}
	for _, row := range f.Rows {
		result += "<div class=\"form-row\">\n"
		for _, item := range row {
			result += item
		}
		result += "</div>\n"

	}
	return result
}
