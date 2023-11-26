package account

import (
	"elichika/protocol/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// this package is for exporting account data from db to jsons and parsing them back.
// this can help with keeping progress even if database is completely changed.
// the jsons will be response data from the servers, which contain the following types:
// - /login/login: login data, contain pretty much everythings, required for account restoration
// - /bootstrap/fetchBootstrap: contain some stuff like challenge_beginner_completed_ids
// - /present/fetch: present box items
// - /trainingTree/fetchTrainingTree: the training tree details for each card (which practice tile is unlocked)
// - delta patch data from every requests
//
// - while it's possible to handle delta patch data, that would be too tedious, plus getting the jsons would be painful.
// - /present/fetch: TODO for now as I don't truly understand that system yet
// - /bootstrap/fetchBootstrap: ditto
// - /trainingTree/fetchTrainingTree: see cars_solver.go

// export to string to write to file or to return to webui
func ExportUser(ctx *gin.Context) string {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	loginData := session.Login()
	text, err := json.Marshal(loginData)
	utils.CheckErr(err)
	return string(text)
}

func ImportUser(ctx *gin.Context, loginJson string, userID int) string {
	loginData := response.Login{}
	loginData.UserModel = new(response.UserModel)
	loginData.UserModel.UserStatus.UserID = userID // userID is not stored in the json response
	err := json.Unmarshal([]byte(loginJson), &loginData)
	if err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := loginJson[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
	}
	utils.CheckErr(err)
	// insert an account based on the login data, not actually inserted into database until finalize is called
	loginData.SetUserID(userID)
	session := userdata.SessionFromImportedLoginData(ctx, &loginData)
	defer session.Close()
	// insert training tree data to make training consistent
	solver := TrainingTreeSolver{}
	solver.LoadUserLogin(&loginData)
	for i := range loginData.UserModel.UserCardByCardID.Objects {
		solver.SolveCard(session, &loginData.UserModel.UserCardByCardID.Objects[i])
	}
	// update term of use and stuff

	session.Finalize("{}", "user_model")

	return "OK"
}
