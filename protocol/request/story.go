package request

type FinishUserStoryMainRequest struct {
	CellId     int32  `json:"cell_id"`
	IsAutoMode bool   `json:"is_auto_mode"`
	MemberId   *int32 `json:"member_id"`
}

type SaveBrowseStoryMainDigestMovieRequest struct {
	PartId int32 `json:"part_id"`
}

type FinishUserStoryLinkageRequest struct {
	CellId     int32 `json:"cell_id"`
	IsAutoMode bool  `json:"is_auto_mode"`
}
