package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"
	"log"

	"xorm.io/xorm"
)

// only load daily theater of this locale
type DailyTheater struct {
	DailyTheaterId int32                `xorm:"pk" json:"daily_theater_id"`
	Title          client.LocalizedText `xorm:"'title'"`
	DetailText     client.LocalizedText `xorm:"'detail_text'"`
	Year           int32                `xorm:"'year'"`
	Month          int32                `xorm:"'month'"`
	Day            int32                `xorm:"'day'"`
	PublishedAt    int64                `xorm:"'published_at'"`
	Members        []int32              `xorm:"-"`
	IsInClient     bool                 `xorm:"-"`
}

func loadDailyTheater(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading DailyTheater")
	gamedata.DailyTheater = make(map[int32]*DailyTheater)
	err := serverdata_db.Table("s_daily_theater").Where("lang = ?", gamedata.Language).Find(&gamedata.DailyTheater)
	utils.CheckErr(err)
	for _, dailyTheater := range gamedata.DailyTheater {
		if gamedata.LastestDailyTheaterId < dailyTheater.DailyTheaterId {
			gamedata.LastestDailyTheaterId = dailyTheater.DailyTheaterId
		}
		err = serverdata_db.Table("s_daily_theater_member").
			Where("lang = ? AND daily_theater_id = ?", gamedata.Language, dailyTheater.DailyTheaterId).
			Cols("member_master_id").Find(&dailyTheater.Members)
		utils.CheckErr(err)
	}

	type DailyTheaterArchiveClient struct {
		Language       string `xorm:"'lang' pk"`
		DailyTheaterId int32  `xorm:"'daily_theater_id' pk"`
		Year           int32  `xorm:"'year'"`
		Month          int32  `xorm:"'month'"`
		Day            int32  `xorm:"'day'"`
		Title          string `xorm:"'title'"`
		PublishedAt    int64  `xorm:"'published_at'"`
	}

	clientDailyTheaterIds := []DailyTheaterArchiveClient{}
	err = masterdata_db.Table("m_daily_theater_archive_client").Where("lang = ?", gamedata.Language).
		Find(&clientDailyTheaterIds)
	utils.CheckErr(err)
	for _, clientDailyTheater := range clientDailyTheaterIds {
		dailyTheater, exist := gamedata.DailyTheater[clientDailyTheater.DailyTheaterId]
		if exist {
			dailyTheater.IsInClient = true
			// this could be reflect but oh well
			if dailyTheater.Year != clientDailyTheater.Year {
				log.Println(fmt.Sprint("ID: ", clientDailyTheater.DailyTheaterId, ". Different year compared to client archive.\nServer: ", dailyTheater.Year, "\nClient: ", clientDailyTheater.Year))
			}
			if dailyTheater.Month != clientDailyTheater.Month {
				log.Println(fmt.Sprint("ID: ", clientDailyTheater.DailyTheaterId, ". Different month compared to client archive.\nServer: ", dailyTheater.Month, "\nClient: ", clientDailyTheater.Month))
			}
			if dailyTheater.Day != clientDailyTheater.Day {
				log.Println(fmt.Sprint("ID: ", clientDailyTheater.DailyTheaterId, ". Different day compared to client archive.\nServer: ", dailyTheater.Day, "\nClient: ", clientDailyTheater.Day))
			}
			if dailyTheater.Title.DotUnderText != clientDailyTheater.Title {
				log.Println(fmt.Sprint("ID: ", clientDailyTheater.DailyTheaterId, ". Different title compared to client archive.\nServer: ", dailyTheater.Title.DotUnderText, "\nClient: ", clientDailyTheater.Title))
			}
			if dailyTheater.PublishedAt != clientDailyTheater.PublishedAt {
				log.Println(fmt.Sprint("ID: ", clientDailyTheater.DailyTheaterId, ". Different published time compared to client archive.\nServer: ", dailyTheater.PublishedAt, "\nClient: ", clientDailyTheater.PublishedAt))
			}
		}
	}
}

func init() {
	addLoadFunc(loadDailyTheater)
}
