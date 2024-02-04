package member_guild

import (
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(member_guild): the logic of this part is wrong or missing
func fetchMemberGuildRankingYear(ctx *gin.Context) {
	fetchMemberGuildRanking(ctx)
}

func init() {
	router.AddHandler("/memberGuild/fetchMemberGuildRankingYear", fetchMemberGuildRankingYear)
}
