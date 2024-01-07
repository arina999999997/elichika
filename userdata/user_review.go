package userdata

import (
	"elichika/utils"
)

func reviewRequestProcessFlowFinalizer(session *Session) {
	for _, userReview := range session.UserModel.UserReviewRequestProcessFlowById.Objects {
		affected, err := session.Db.Table("u_review_request_process_flow").
			Where("user_id = ? AND review_request_id = ?",
				session.UserStatus.UserId, userReview.ReviewRequestId).
			AllCols().Update(userReview)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_review_request_process_flow").
				Insert(userReview)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(reviewRequestProcessFlowFinalizer)
	addGenericTableFieldPopulator("u_review_request_process_flow", "UserReviewRequestProcessFlowById")
}
