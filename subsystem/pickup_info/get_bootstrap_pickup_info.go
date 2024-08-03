package pickup_info

import (
	"elichika/client"
	"elichika/generic"
	"elichika/subsystem/event"
	"elichika/userdata"
)

// TODO(pickup_info): Fill this with more relevant event items
// and maybe put the gacha things into a database or something
func GetBootstrapPickupInfo(session *userdata.Session) client.BootstrapPickupInfo {
	resp := client.BootstrapPickupInfo{}
	resp.ActiveEvent = event.GetActiveEventPickup(session)
	// birthday scouting
	// resp.AppealGachas.Append(client.TextureStruktur{V: generic.NewNullable("'-K")})
	// muse festival party
	resp.AppealGachas.Append(client.TextureStruktur{V: generic.NewNullable("'/&")})
	// aqours festival party
	resp.AppealGachas.Append(client.TextureStruktur{V: generic.NewNullable("z7j")})
	// niji festival party
	resp.AppealGachas.Append(client.TextureStruktur{V: generic.NewNullable("Q%T")})
	return resp
}
