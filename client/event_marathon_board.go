package client

import (
	"elichika/generic"
)

type EventMarathonBoard struct {
	BoardThingMasterRows generic.List[EventMarathonBoardMemorialThingsMasterRow] `json:"board_thing_master_rows"`
	IsEffect             bool                                                    `json:"is_effect"`
	BoardBaseImagePath   TextureStruktur                                         `json:"board_base_image_path"`
	BoardDecoImagePath   TextureStruktur                                         `json:"board_deco_image_path"`
}
