package enum

const (
	MinuteSecondCount int32 = 60
	HourMinuteCount   int32 = 60
	HourSecondCount   int32 = HourMinuteCount * MinuteSecondCount
	DayHourCount      int32 = 24
	DaySecondCount    int32 = DayHourCount * HourSecondCount
	WeekDayCount      int32 = 7
	WeekSecondCount   int32 = WeekDayCount * DaySecondCount
)
