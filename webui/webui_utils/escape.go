package webui_utils

import (
	"elichika/utils"

	"encoding/json"
)

// return a js string with proper escaping
// note that this string should be used inside a js file or inside <script></script>
// for using it inside an html element (i.e. as attributes, use html.EscapeString)
func JsString(s string) string {
	bytes, err := json.Marshal(s) // js and json pretty much use the same rule
	utils.CheckErr(err)
	return string(bytes)
}
