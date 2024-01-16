package request

type StartupRequest struct {
	Mask                        string `json:"mask"`
	ResemaraDetectionIdentifier string `json:"resemara_detection_identifier"`
	TimeDifference              int32  `json:"time_difference"`
	RecaptchaToken              string `json:"recaptcha_token"`
}
