package request

type FinishUserStoryMainRequest struct {
	CellID     int  `json:"cell_id"`
	IsAutoMode bool `json:"is_auto_mode"`
	MemberID   *int `json:"member_id"`
}

type SaveBrowseStoryMainDigestMovieRequest struct {
	PartID int `json:"part_id"`
}

type FinishUserStoryLinkageRequest struct {
	CellID     int  `json:"cell_id"`
	IsAutoMode bool `json:"is_auto_mode"`
}
