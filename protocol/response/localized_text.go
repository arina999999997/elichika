package response

// no idea why it's named this, since it's used for things that aren't localized as well
// maybe localized is used in the sense of distance, not country
type LocalizedText struct {
	DotUnderText string `json:"dot_under_text"`
}
