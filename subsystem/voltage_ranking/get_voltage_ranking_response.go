package voltage_ranking

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/subsystem/user_social"
	"elichika/userdata"
)

func GetVoltageRankingResponse(session *userdata.Session, liveDifficultyId int32) response.GetVoltageRankingResponse {
	records := GetRankingByLiveDifficultyId(session, liveDifficultyId).GetRange(1, VoltageRankingLimit)
	resp := response.GetVoltageRankingResponse{}
	for i, record := range records {
		resp.VoltageRankingCells.Append(client.VoltageRankingCell{
			Order:              int32(i + 1), // no tie handling
			VoltagePoint:       record.Score,
			VoltageRankingUser: user_social.GetVoltageRankingUser(session, record.Id),
		})
	}
	return resp
}
