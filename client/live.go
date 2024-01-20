package client

import (
	"elichika/generic"

	"encoding/json"
	"reflect"
)

type Live struct {
	LiveId          int64                           `xorm:"pk" json:"live_id"`
	LiveType        int32                           `json:"live_type"`
	DeckId          int32                           `json:"deck_id"`
	LiveStage       LiveStage                       `xorm:"json" json:"live_stage"`
	LivePartnerCard generic.Nullable[OtherUserCard] `xorm:"json" json:"live_partner_card"`
	IsPartnerFriend bool                            `json:"is_partner_friend"`
	CellId          generic.Nullable[int32]         `xorm:"json" json:"cell_id"`
	TowerLive       generic.Nullable[TowerLive]     `xorm:"json" json:"tower_live"`
}

// special behavior to not marshal tower live as null if it's empty
func (l Live) MarshalJSON() ([]byte, error) {
	bytes := []byte("{")
	rType := reflect.TypeOf(l)
	isFirst := true
	for i := 0; i < rType.NumField(); i++ {
		rField := rType.Field(i)
		if (rField.Name == "TowerLive") && (!l.TowerLive.HasValue) {
			continue
		}
		key := rField.Tag.Get("json")
		if key == "-" {
			continue
		} else if key == "" {
			panic("empty key")
		}
		if isFirst {
			isFirst = false
		} else {
			bytes = append(bytes, []byte(",")...)
		}
		bytes = append(bytes, []byte("\"")...)
		bytes = append(bytes, []byte(key)...)
		bytes = append(bytes, []byte("\":")...)
		fieldBytes, err := json.Marshal(reflect.ValueOf(l).Field(i).Interface())
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, fieldBytes...)
	}

	bytes = append(bytes, []byte("}")...)
	return bytes, nil
}
