package client

type UserGdprConsentedInfo struct {
	HasConsentedAdPurposeOfUse bool `json:"has_consented_ad_purpose_of_use"`
	HasConsentedCrashReport    bool `json:"has_consented_crash_report"`
}
