package user

type Response struct {
	Response *string `json:"response"`
	Error    *string `json:"error"`
}
