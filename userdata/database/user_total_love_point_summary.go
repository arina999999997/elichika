package database

type UserTotalLovePointSummary struct {
	UserId     int32 `xorm:"pk 'user_id'"`
	All        int32 `xorm:"'condition_1'"`
	Muse       int32 `xorm:"'condition_2'"`
	Aqours     int32 `xorm:"'condition_3'"`
	Niji       int32 `xorm:"'condition_4'"`
	Printemps  int32 `xorm:"'condition_35'"`
	Bibi       int32 `xorm:"'condition_36'"`
	LilyWhite  int32 `xorm:"'condition_37'"`
	Cyaron     int32 `xorm:"'condition_38'"`
	Azalea     int32 `xorm:"'condition_39'"`
	GuiltyKiss int32 `xorm:"'condition_40'"`
	Azuna      int32 `xorm:"'condition_41'"`
	DiverDiva  int32 `xorm:"'condition_42'"`
	Qu4rtz     int32 `xorm:"'condition_43'"`
	NijiUnit4  int32 `xorm:"'condition_44'"`
}

func init() {
	AddTable("u_total_love_point_summary", UserTotalLovePointSummary{})
}
