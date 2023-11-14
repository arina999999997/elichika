package response

// it get too silly putting the response struct in handler, especially when member_guild is very complicated
// the naming is derived from decompiled code
// int int64 is derived from reading the code, but there might be mistake
// and they don't always make sense

type MemberGuildRankingOneTermCell struct {
	Order          int `json:"order"`
	TotalPoint     int `json:"total_point"`
	MemberMasterID int `json:"member_master_id"`
}

type MemberGuildRankingOneTerm struct {
	MemberGuildID int                             `json:"member_guild_id"` // this is not the ID of the member, this is what week this ranking is from
	StartAt       int64                           `json:"start_at"`
	EndAt         int64                           `json:"end_at"`
	Channels      []MemberGuildRankingOneTermCell `json:"channels"`
}
type MemberGuildRanking struct {
	ViewYear               int                         `json:"view_year"`
	NextYear               int                         `json:"next_year"` // can be missing
	PreviousYear           int                         `json:"previous_year"`
	MemberGuildRankingList []MemberGuildRankingOneTerm `json:"member_guild_ranking_list"`
}

type MemberGuildUserRankingUserData struct {
	UserID                 int           `json:"user_id"`
	UserName               LocalizedText `json:"user_name"`
	UserRank               int           `json:"user_rank"`
	CardMasterID           int           `json:"card_master_id"`
	Level                  int           `json:"level"`
	IsAwakening            bool          `json:"is_awakening"`
	IsAllTrainingActivated bool          `json:"is_all_training_activated"`
	EmblemMasterID         int           `json:"emblem_master_id"`
}

type MemberGuildUserRankingCell struct {
	Order                          int64                          `json:"order"`
	TotalPoint                     int                            `json:"total_point"`
	MemberGuildUserRankingUserData MemberGuildUserRankingUserData `json:"member_guild_user_ranking_user_data"`
}

type MemberGuildUserRankingBorderInfo struct {
	RankingOrderPoint int   `json:"ranking_border_point"` // probably the ID
	UpperRank         int   `json:"upper_rank"`
	LowerRank         int64 `json:"lower_rank"`
	DisplayOrder      int   `json:"display_order"`
}

type MemberGuildUserRanking struct {
	MemberGuildID  int                                `json:"member_guild_id"` // this is not the ID of the member, this is what week this ranking is from
	TopRanking     []MemberGuildUserRankingCell       `json:"top_ranking"`
	MyRanking      []MemberGuildUserRankingCell       `json:"my_ranking"`
	RankingBorders []MemberGuildUserRankingBorderInfo `json:"ranking_borders"`
}

type FetchMemberGuildRankingResponse struct {
	MemberGuildRanking         MemberGuildRanking       `json:"member_guild_ranking"`
	MemberGuildUserRankingList []MemberGuildUserRanking `json:"member_guild_user_ranking_list"`
}
