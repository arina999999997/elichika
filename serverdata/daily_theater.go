package serverdata

import (
	"elichika/config"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"xorm.io/xorm"
)

type DailyTheater struct {
	Language       string `xorm:"pk 'lang'" json:"language"`
	DailyTheaterId int32  `xorm:"pk" json:"daily_theater_id"`
	Year           int32  `json:"year"`
	Month          int32  `json:"month"`
	Day            int32  `json:"day"`
	Title          string `json:"title"`
	DetailText     string `json:"detail_text"`
	PublishedAt    int64  `xorm:"published_at"`
}
type DailyTheaterMember struct {
	Language       string `xorm:"pk 'lang'" json:"language"`
	DailyTheaterId int32  `xorm:"pk" json:"daily_theater_id"`
	MemberMasterId int32  `xorm:"pk"`
}

// Daily theater is stored and loaded from serverdata.db
// The data for it is hosted in the assets database, so technically it's stored twice
// The format for storing daily theater:
// - store daily theaters inside /assets/daily_theater
// - one daily theater per json file:
//   - multiple languages requires different files
//   - the file name has no effect, only the info inside the file do.
// - all dirs will be recursively explored
//
// The daily theater member will be extracted from the text itself:
// - Every <:th_ch0xxx/> tags will be extracted, and the member is saved
// - The matching is done with regex, limited a bit.

func InitializeDailyTheater(session *xorm.Session) {
	asiaTokyo, _ := time.LoadLocation("Asia/Tokyo")
	memberRegex := regexp.MustCompile(`<:th_ch0[0-2][0-1][0-9]/>`)
	filepath.Walk(config.AssetPath+"daily_theater", func(path string, info os.FileInfo, err error) error {
		utils.CheckErr(err)
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		fmt.Printf("Parsing daily theater file: %s\n", path)
		text := utils.ReadAllText(path)
		dailyTheater := DailyTheater{}
		err = json.Unmarshal([]byte(text), &dailyTheater)
		utils.CheckErr(err)
		dailyTheater.PublishedAt = time.Date(int(dailyTheater.Year), time.Month(dailyTheater.Month), int(dailyTheater.Day),
			0, 0, 0, 0, asiaTokyo).Unix()

		_, err = session.Table("s_daily_theater").Insert(dailyTheater)
		utils.CheckErr(err)
		memberMatches := memberRegex.FindAllString(dailyTheater.DetailText, -1)
		members := map[int32]bool{}
		for _, member := range memberMatches {
			member = member[8:11]
			memberId, err := strconv.Atoi(member)
			utils.CheckErr(err)
			if memberId == 0 {
				continue
			}
			members[int32(memberId)] = true
		}
		for memberId := range members {
			session.Table("s_daily_theater_member").Insert(DailyTheaterMember{
				Language:       dailyTheater.Language,
				DailyTheaterId: dailyTheater.DailyTheaterId,
				MemberMasterId: memberId,
			})
		}
		return nil
	})
}

func init() {
	addTable("s_daily_theater", DailyTheater{}, InitializeDailyTheater)
	addTable("s_daily_theater_member", DailyTheaterMember{}, nil)
}
