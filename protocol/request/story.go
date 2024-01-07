package request

type FinishUserStoryMainRequest struct {
	CellId     int  `json:"cell_id"`
	IsAutoMode bool `json:"is_auto_mode"`
	MemberId   *int `json:"member_id"`
}

type SaveBrowseStoryMainDigestMovieRequest struct {
	PartId int `json:"part_id"`
}

type FinishUserStoryLinkageRequest struct {
	CellId     int  `json:"cell_id"`
	IsAutoMode bool `json:"is_auto_mode"`
}
