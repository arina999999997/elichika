package serverdata

import (
	"elichika/config"
	"elichika/parser"
	"elichika/utils"

	"os"
	"strconv"

	"xorm.io/xorm"
)

type EventAvailable struct {
	EventId int32 `xorm:"pk 'event_id'"`
	Order   int32 `xorm:"unique 'order'"`
}

func init() {
	addTable("s_event_available", EventAvailable{}, initEventAvailable)
}

func initEventAvailable(session *xorm.Session) {
	eventOrders := []int32{}
	parser.ParseJson(config.AssetPath+"event/event_order.json", &eventOrders)
	eventOrderMap := map[int32]int32{}
	for order, eventId := range eventOrders {
		eventOrderMap[eventId] = int32(order + 1)
	}
	events := []EventAvailable{}
	{
		entries, err := os.ReadDir(config.AssetPath + "event/marathon")
		utils.CheckErr(err)
		for _, entry := range entries {
			eventId, err := strconv.Atoi(entry.Name())
			utils.CheckErr(err)
			events = append(events, EventAvailable{
				EventId: int32(eventId),
				Order:   eventOrderMap[int32(eventId)],
			})
		}
	}
	{
		entries, err := os.ReadDir(config.AssetPath + "event/mining")
		utils.CheckErr(err)
		for _, entry := range entries {
			eventId, err := strconv.Atoi(entry.Name())
			utils.CheckErr(err)
			events = append(events, EventAvailable{
				EventId: int32(eventId),
				Order:   eventOrderMap[int32(eventId)],
			})
		}
	}
	_, err := session.Table("s_event_available").Insert(events)
	utils.CheckErr(err)
}
