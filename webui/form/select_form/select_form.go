package select_form

import (
	"elichika/webui/webui_utils"

	"fmt"
	"html"
	"strings"
)

// a form where we can select between values
// look something like this:
// - <DataLabel>: <Textbox>
// - Various select cases which are actually button that change the textbox's input:
//   - Clicking on the button will set the textbox value
//
// - We can manually enable textbox manual input if we want to allow unaccounted for input
type SelectForm struct {
	FormId    string
	DataLabel string
	DataId    string
	Keys      []string
	Texts     []string
}

func (f *SelectForm) AddLinebreak() {
	f.Keys = append(f.Keys, "line_break")
	f.Texts = append(f.Texts, "line_break")
}

func (f *SelectForm) AddOption(value, text string) {
	f.Keys = append(f.Keys, value)
	f.Texts = append(f.Texts, text)
}

func jsEscape(s string) string {
	return html.EscapeString(strings.ReplaceAll(s, `\`, `\\`))
}

func (f *SelectForm) GetHTML() string {
	result := `<form id="` + f.FormId + `" method="POST" enctype="multipart/form-data">`
	result += "\n"
	result += `<div><label>` + f.DataLabel + `:</label> <input type="text" id="` + f.DataId + `" name="` + f.DataId + `" autofocus`
	result += "/>"
	result += "\n"
	result += `</form><br>`
	result += "<div class=\"form-row\">\n"
	for i, value := range f.Keys {
		result += "\n"
		if value == "line_break" {
			result += "</div>\n"
			result += "<div class=\"form-row\">\n"
			continue
		}
		script := fmt.Sprintf(`set_form_data(this, %s, %s)`, webui_utils.JsString(f.DataId), webui_utils.JsString(value))
		result += fmt.Sprintf(`<button class="form-option form-button-option" type="button" onclick="%s">`, html.EscapeString(script))
		fmt.Println(f.Texts[i])
		result += html.EscapeString(f.Texts[i])
		result += `</button>`
	}
	return result
}
