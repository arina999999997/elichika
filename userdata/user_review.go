package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func reviewRequestProcessFlowFinalizer(session *Session) {
	for id, userReview := range session.UserModel.UserReviewRequestProcessFlowById.Map {
		affected, err := session.Db.Table("u_review_request_process_flow").
			Where("user_id = ? AND review_request_id = ?",
				session.UserId, id).
			AllCols().Update(userReview)
		utils.CheckErr(err)
		if affected == 0 {
			type Temp struct {
				ReviewRequestId              int64                               `xorm:"pk 'review_request_id'"`
				UserReviewRequestProcessFlow client.UserReviewRequestProcessFlow `xorm:"extends"`
			}
			temp := Temp{
				ReviewRequestId:              id,
				UserReviewRequestProcessFlow: *userReview,
			}
			GenericDatabaseInsert(session, "u_review_request_process_flow", temp)
		}
	}
}

func init() {
	AddContentFinalizer(reviewRequestProcessFlowFinalizer)
}
