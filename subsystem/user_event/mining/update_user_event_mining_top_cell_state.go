package mining

import (
	"elichika/userdata"
	"elichika/userdata/database"
)

// note that it's impossible to unlike a cell using the client, so we only need to insert here
func UpdateUserEventMiningTopCellState(session *userdata.Session, eventId, thumbnailCellId int32) {
	userdata.GenericDatabaseInsert(session, "u_event_mining_top_cell_state", database.UserEventMiningTopCellState{
		EventId:         eventId,
		ThumbnailCellId: thumbnailCellId,
		IsLike:          true,
	})
}
