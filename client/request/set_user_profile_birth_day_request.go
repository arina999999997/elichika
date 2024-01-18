package request

type SetUserProfileBirthDayRequest struct {
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}
