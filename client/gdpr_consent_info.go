package client

type GdprConsentInfo struct {
	GdprType     int32  `json:"gdpr_type" enum:"GdprConsentType"`
	HasConsented bool   `json:"has_consented"`
	AdIdentifier string `json:"ad_identifier"`
}
