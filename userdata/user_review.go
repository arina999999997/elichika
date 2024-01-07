package userdata

import (
	"elichika/utils"
)

func reviewRequestProcessFlowFinalizer(session *Session) {
	for _, userReview := range session.UserModel.UserReviewRequestProcessFlowById.Objects {
		affected, err := session.Db.Table("u_review_request_process_flow").
			Where("user_id = ? AND review_request_id = ?",
				session.UserId, userReview.ReviewRequestId).
			AllCols().Update(userReview)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_review_request_process_flow", userReview)
		}
	}
}

func init() {
	addFinalizer(reviewRequestProcessFlowFinalizer)
	addGenericTableFieldPopulator("u_review_request_process_flow", "UserReviewRequestProcessFlowById")
}
