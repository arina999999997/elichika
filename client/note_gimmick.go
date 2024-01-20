package client

type NoteGimmick struct {
	UniqId          int32 `json:"uniq_id"`
	Id              int32 `json:"id"`
	NoteGimmickType int32 `json:"note_gimmick_type" enum:"NoteGimmickType"`
	Arg1            int32 `json:"arg_1"`
	Arg2            int32 `json:"arg_2"`
	EffectMId       int32 `json:"effect_m_id"`
	IconType        int32 `json:"icon_type" enum:"NoteGimmickIconType"`
}

func (ng *NoteGimmick) IsSame(other *NoteGimmick) bool {
	same := true
	same = same && (ng.UniqId == other.UniqId)
	same = same && (ng.NoteGimmickType == other.NoteGimmickType)
	same = same && (ng.Arg1 == other.Arg1)
	same = same && (ng.Arg2 == other.Arg2)
	same = same && (ng.EffectMId == other.EffectMId)
	if !same {
		return false
	}
	if ng.IconType == other.IconType {
		return true
	}
	if ng.IconType == 5 && other.IconType == 25 { // there was a db update that change this
		return true
	}
	if ng.IconType == 8 && other.IconType == 9 { // there was a db update that change this
		return true
	}
	return false
}
