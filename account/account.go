package account

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// this package is for exporting account data from db to jsons and parsing them back.
// this can help with keeping progress even if database is completely changed.
// the jsons will be response data from the servers, which contain is the response of:
// - /login/login: login data, contain pretty much everythings, required for account restoration
// there can be other info stored in other network endpoint, but they are not taken care of for now.
// it's better to just use the userdata.db format as it contain everything and is more lightweight.
// TODO(extra): Maybe implement a system that read from various json and accept delta patching.
// That won't be too helpful as an exporting format, but it can help with people who have recorded network data
// (which is admitedly very small).

// export login data to json string to write to file or to return to webui
func ExportLoginJson(session *userdata.Session) []byte {
	loginData := session.Login()
	text, err := json.Marshal(loginData)
	utils.CheckErr(err)
	return text
}

func ImportUserJson(ctx *gin.Context, loginJson []byte) string {
	loginData := response.LoginResponse{}
	loginData.UserModel = new(client.UserModel)

	err := json.Unmarshal([]byte(loginJson), &loginData)
	if err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := loginJson[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
	}
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	session.ImportLoginData(ctx, &loginData)

	// insert training tree data to make training consistent
	solver := TrainingTreeSolver{}
	solver.LoadUserLogin(&loginData)
	for i := range loginData.UserModel.UserCardByCardId.Map {
		solver.SolveCard(session, *loginData.UserModel.UserCardByCardId.Map[i])
	}
	// update term of use and stuff

	// actually populate the tracking field so we can import items
	user_content.PopulateGenericContentDiffFromUserModel(session)
	session.Finalize()
	return "Imported json data"
}
