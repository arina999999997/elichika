package user

import (
	"elichika/account"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"
	"elichika/webui/webui_utils"

	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func transferAccount(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)
	session.DeleteUserGameData()

	form, _ := ctx.MultipartForm()
	fileHeader := form.File["file"][0]
	ext := path.Ext(fileHeader.Filename)
	resp := webui_utils.Response{
		Response: new(string),
	}
	if (ext != ".json") && (ext != ".db") {
		*resp.Response = "Must be .db or .json"
	} else {
		file, err := fileHeader.Open()
		utils.CheckErr(err)
		bytes := make([]byte, fileHeader.Size)
		length, err := file.Read(bytes)
		utils.CheckErr(err)
		if int64(length) != fileHeader.Size {
			panic("error reading file")
		}
		switch ext {
		case ".json":
			*resp.Response = account.ImportUserJson(ctx, bytes)
		case ".db":
			resp.Response, resp.Error = session.ImportDatabaseData(ctx, bytes)
		}
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func exportJson(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	content := account.ExportLoginJson(session)

	ctx.Header("Content-Disposition", fmt.Sprint("attachment; filename=login_", session.UserId, ".json"))
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Accept-Length", fmt.Sprint(len(content)))
	ctx.Writer.Write(content)
}

func exportDb(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	content := session.ExportDb()
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", fmt.Sprint("attachment; filename=backup_", session.UserId, ".db"))
	ctx.Header("Content-Type", "application/octet-stream")
	// ctx.Header("Accept-Length", fmt.Sprint(len(content)))
	ctx.Writer.Write(content)
}

func transferForm(ctx *gin.Context) {
	form :=
		`<div><label>Import or export account. Note that importing account WILL RESET CURRENT PROGRESS, so you might want to create a backup through export first.</label></div>
<div><label>Exporting to .db will export all the data stored by server, but the data will only works for (this and future versions of) elichika, and servers that support this format.</label></div>
<div><label>Exporting to .json will only export login data as received by the game, which contain most relevant data but not everything. This data is canonical and should be supported by all server implementation, if they do support account import.</label></div>
<div><label>Using japanese server data on global client can lead to issues, but you can import GL account to JP server and it should still work.</label></div>
<div><label>For importing account from pcap, extract it to a json first, learn about it <a href="https://github.com/arina999999997/elichika/tree/master/docs">here.</a></label></div>
<form id="transfer_form">
<div><input type="button" value="Export to .db" onclick="submit_form(null, 'export_db', true)"></div>
<div><input type="button" value="Export to .json" onclick="submit_form(null, 'export_json', true)"></div>
<div>
<input type="file" name="file">
<input type="button" value="Import account" onclick="if (confirm('This will overwrite existing data, import account?')) submit_form('transfer_form', 'transfer')">
</div>
</form>
`
	ctx.HTML(http.StatusOK, "logged_in_user.html", gin.H{
		"body": form,
	})
}

func init() {
	addFeature("Import / Export account", "transfer")
	router.AddHandler("/webui/user", "GET", "/transfer", transferForm)
	router.AddHandler("/webui/user", "POST", "/export_json", exportJson)
	router.AddHandler("/webui/user", "POST", "/export_db", exportDb)
	router.AddHandler("/webui/user", "POST", "/transfer", transferAccount)
}
