package client

import (
	"elichika/generic"
)

type ResumeFinishInfo struct {
	CachedJudgeResult generic.Dictionary[int32, int32] `json:"cached_judge_result" enum:"JudgeType"`
}
