package client

type LiveNoteSetting struct {
	Id                  int32 `json:"id"`
	CallTime            int32 `json:"call_time"`
	NoteType            int32 `json:"note_type" enum:""`
	NotePosition        int32 `json:"note_position"`
	GimmickId           int32 `json:"gimmick_id"`
	NoteAction          int32 `json:"note_action" enum:""`
	WaveId              int32 `json:"wave_id"`
	NoteRandomDropColor int32 `json:"note_random_drop_color" enum:""` // this control drop and can change between run
	AutoJudgeType       int32 `json:"auto_judge_type" enum:""`        // this is controlled by config
}

func (lns *LiveNoteSetting) IsSame(other *LiveNoteSetting) bool {
	same := true
	same = same && (lns.Id == other.Id)
	same = same && (lns.CallTime == other.CallTime)
	same = same && (lns.NoteType == other.NoteType)
	same = same && (lns.NotePosition == other.NotePosition)
	same = same && (lns.GimmickId == other.GimmickId)
	same = same && (lns.NoteAction == other.NoteAction)
	same = same && (lns.WaveId == other.WaveId)
	return same
}
